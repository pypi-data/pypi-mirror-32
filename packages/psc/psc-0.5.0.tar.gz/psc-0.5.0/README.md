# psc: Parameter Store Client - an EC2 SSM Parameter Store CLI helper

## Properly managing secrets and configs

Secrets should not be stored unencrypted, and their exposure should be limited.
Therefore it make sense to use a proper encrypted store to manage our secrets.

We built this with the following goals in mind:
- We must have fine-grain control on who has access to which secret in any given environment
- Secrets must be stored very securely (including the master key problem, aka the [turtles all the way down](https://www.youtube.com/watch?v=OUSvv2maMYI) problem)
- API calls done to read or write secrets must be auditable.
- Access to secrets is critical, and must be highly available
- Adding security for secrets should not add a burden on developers for local config

We want to keep things simple for local development, our localhost db URL isn't a critical security entry point and using it shouldn't add any constraint.

Because our platform run on AWS, we choose to leverage one of the component of AWS EC2 Simple Systems Manager (SSM) service, namely, the [Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-paramstore.html).
The Parameter Store lets you leverage KMS strong encryption capabitlies, control who can read and write twith finegrained IAM policies. In short, it gives you a fully-managed, secure, resilient and cheap secret store, solving most of our problems. This CLI is meant to be a convenient way to set secrets in the Parameter Store and to consume secrets form it.
Parameter Store itself don't provide the access audit logs, but as soon as you enable Cloudtrail, you will get these access logs.

In order to make it easy to use, this secret cli has only a few commands:
- set let you store a secret
- get let you decrypt a secret
- delete
- decrypt-env search for specific pattern in your env vars and generate the `export` commands to update your env with decrypted secrets, if any match is found
- ls (and ll) let you list your secrets


## Avoid any burden on local development

The idea is that the decrypted secret will be available for the code from an env var, like many other configs.
In dev, the setup could be like `export DB_URL="mongodb://localhost/myservice"` and you use it in your code with `process.env.DB_URL`.

Now, we want to be able to encrypt the production value of this db connection string, then at runtime decrypt it and populate theat env var.

The idea is that the `decrypt-env` command will look for env vars with the following pattern: _VAR_NAME_=SECRET _encrypted_secret_key_ . More specifically, it will check for this pattern: `^(\w+)\=SECRET (.+)$`
For any mathcing entry, it will ask the Parameter Store if a secret is stored with the key put in that env var. If nothing is found, the environment variable is left untouched. if a secret has been found, the scritpt generate an `export VAR_NAME=decrypted value` command, that your shell can source before to start your service. This will overwrite the envrionment var with the decrypted secret, exposing this value to your service.


## Installation

PSC is a Python library, therefore Python and pip are required. Then it's as easy as:

    pip install psc


## Usage overview

**DEV SETUP** : `export DB_URL="mongodb://localhost/myservice"`  
**STAGING SETUP** : `export DB_URL="SECRET staging.myservice.db_url"`

In dev, the 'secret' is available in the env at any time, easy for daily work.

For staging the secret owner/admin can set the secret with
`psc set staging.myservice.db_url "mongodb://localhost/myservice" --description "my service connection string to store critical data" --debug`

In staging (or prod), we'll use `psc decrypt-env` to substitute `SECRET staging.myservice.db_url` with the actual secret references in SSM, stored just before:

```sh
$ echo $DB_URL; source <(psc decrypt-env) ; node -e "console.log(process.env.DB_URL);"
SECRET staging.myservice.db_url
mongodb://localhost/myservice
```
Obviously, in the real world, we're not going to just echo the secret (we should not, it's the decrytped value that would appear in the logs!), but this shows we correctly subsititued the id with the decrypted value before to launch the node server (or any other program you'd want to run).


## How to use it

I assume you have the aws-cli already setup. If it's not yet done, please check the [AWS CLI documentation](https://docs.aws.amazon.com/cli/latest/userguide/installing.html).
If you have roles or profiles setup in your AWS configuration, you can use them. This is quite useful to manage secret in other accounts, or only use temporary credentials.


#### Store your secrets in the Parameter Store

```sh
psc set staging.test.VAR1 "secret 1"
psc set staging.test.VAR2 "secret 2"
psc set staging.test.VAR4 "secret 4"
psc set staging.myservice.db_url "mongodb://admin:secret@localhost/myservice" --description "my service connection string to store critical data" --debug
```

You can check it's correctly saved with the `get` command:

`psc get "staging.myservice.db_url"`


#### Setup your app to decrypt and inject these secrets at launch

Here are some secrets we define in the provisionning of our service (either via cloudinit if you're running EC2, in the task definition for ECS, in a Kubernetes Config... )  
Note that VAR3 hasn't been defined in our previous step.

```sh
export VAR1="SECRET staging.test.VAR1"
export VAR2="SECRET staging.test.VAR2"
export VAR3="SECRET staging.test.VAR3"
export VAR4="SECRET staging.test.VAR4"
export LOG_LEVEL=debug
export DB_URL="SECRET staging.myservice.db_url"
```

Now, you only need to source the `psc decrypt-env` output before to start your server app:

```sh
echo $DB_URL;echo "----- DECRYPT ----"; source <(psc decrypt-env) ; node -e '["VAR1", "VAR2", "VAR3", "VAR4", "LOG_LEVEL", "DB_URL"].forEach((s) => console.log(`${s}:${process.env[s]}`));'
mongodb://localhost/myservice
----- DECRYPT ----
VAR1:secret 1
VAR2:secret 2
VAR3:SECRET staging.test.VAR3
VAR4:secret 4
LOG_LEVEL:debug
DB_URL:mongodb://admin:secret@localhost/myservice
```

Depending on your use case, you load the update in various ways, like `source <(psc decrypt-env)` or `eval $(psc decrypt-env)`.

## Restrictions

The name of the parameter you store in the Parameter Store (ie the name to access your stored secret) is scoped by account. Let's say in us-east-1 you'll use many micro-services for a given aws account, they might have different DB_URL connection strings. but the parameter name `DB_URL` has a unique scope for this region and this account.

In order to work around this, we use a naming convention like this `{stack}.{service}.{name}`. For instance, it could be `staging.billing.DB_URL` (Obviously you can use whatever pattern you want, or use underscores. The point is the key must be unique.)

Finally, the Parameter Store has some other subtle constraints:
- A parameter name must be unique within your AWS account.
- Parameter names are case-sensitive.
- A parameter name can't be prefixed with `aws` or `ssm` (case-insensitive). For example, `awsTestParameter` or `SSM-testparameter` will fail with an exception.
- Parameter names can only include the following symbols and letters: `a-zA-Z0-9_.-`
- A parameter value cannot be longer than 4096 characters (like for direct KMS API calls).
- there is a limit of 1000 parameters per account (it's likely that it's a soft-limit like for many other AWS services, but we haven't yet checked if we can request a limit increase for this. If you plan to use it so widely, please verify this beforehand)

You can read the full list in the [official documentation](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-paramstore-params.html) and the [Parameter Store limits](http://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits_ssm).


## Security limitations, pros and cons.

This is not a perfect solution for a variety of reasons. For instance the decrypted value is kept in memory for a while and we don't purge it. The security boundaries are set by the IAM role, which might be a challenge if you want to segregate secrets in a multi-tenant setup running various tenants containers on the same host.

On the other hand the goal was to be able to get many benefits with a minimal cost:
- Fully managed solution, no single point of failure. You could host your own Vault solution, but that would be a new critical piece to manage.
- 0 burden and 0 change for local devs. The dev setup and flow is still the same as a before adding secret encryption.
- Strong encryption. You can check all the AWS KMS documentation available, but the bottom line is that you benefit from a strong encryption and a strong key management. Managing keys is so hard to do correctly, we are happy to not have to do it!
- secret access logs. For forensics and compliance requirements, if you enable Cloudtrail, you'll get all the details of the who/what/when the secrets are read and changed.

If you want a more ephemeral life for your decrypted secrets, you could ignore the `psc decrypt-env` command and just call `psc get...` from your code.


## Available Commands

Make sure you call `psc --help` to list all the commands, then `psc COMMAND --help` to get the details about each specific command you'd like.

For instance:
- `psc set --help` to learn more about how to store a secret from a string
- `psc set-file --help` to learn more about how to store a secret a file
- `psc get --help` to learn more about how to decrypt a single secret
- `psc delete --help` to check how to delete a secret
- `psc ls --help` to view the few options to list your stored secrets
- `psc decrypt-env --help` to list the parameters for swapping the secrets IDs and their decrypted values in your env
...
