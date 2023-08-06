import time

import boto3
from botocore.errorfactory import ClientError

from ..component import Component


class IAMRole(Component):

    def __init__(self, name, trust_document, **options):
        self.name = name
        self.trust_document = trust_document
        self.attach_policies = options.get('attach_policies', [])
        super().__init__(**options)
        self.iam = boto3.client('iam')

    def __str__(self):
        return f"IAM role {self.name}"

    def is_setup(self):
        try:
            self.iam.get_role(RoleName=self.name)
            return True
        except ClientError:
            return False

    def set_it_up(self):
        self.iam.create_role(RoleName=self.name, AssumeRolePolicyDocument=self.trust_document)
        for policy_arn in self.attach_policies:
            self.iam.attach_role_policy(RoleName=self.name, PolicyArn=policy_arn)

    def tear_it_down(self):
        for policy in self.iam.list_attached_role_policies(RoleName=self.name)['AttachedPolicies']:
                self.iam.detach_role_policy(RoleName=self.name, PolicyArn=policy['PolicyArn'])

        self.iam.delete_role(RoleName=self.name)


class RoleInlinePolicy(Component):

    def __init__(self, role_name, name, policy_document, **options):
        self.role_name = role_name
        self.name = name
        self.policy_document = policy_document
        super().__init__(**options)
        self.iam = boto3.client('iam')

    def __str__(self):
        return f"IAM role {self.role_name} inline policy {self.name}"

    def is_setup(self):
        try:
            return self.name in self.iam.list_role_policies(RoleName=self.role_name)['PolicyNames']
        except ClientError:
            return False

    def set_it_up(self):
        self.iam.put_role_policy(RoleName=self.name, PolicyName=self.name, PolicyDocument=self.policy_document)

    def tear_it_down(self):
        self.iam.delete_role_policy(RoleName=self.role_name, PolicyName=self.name)


class ServiceLinkedRole(Component):
    def __init__(self, name=None, aws_service_name=None, **options):
        self.name = name
        self.aws_service_name = aws_service_name
        super().__init__(**options)
        self.iam = boto3.client('iam')

    def __str__(self):
        name = self.name or f"for {self.aws_service_name}"
        return f"IAM service-linked role {name}"

    def is_setup(self):
        try:
            self.iam.get_role(RoleName=self.name)
            return True
        except ClientError:
            return False

    def set_it_up(self):
        self.iam.create_service_linked_role(AWSServiceName=self.aws_service_name)

    def tear_it_down(self):
        self.iam.delete_service_linked_role(RoleName=self.name)
        while True:
            try:
                time.sleep(1)
                self.iam.get_role(RoleName=self.name)
            except ClientError:
                return


class Policy(Component):

    def __init__(self, name, **options):
        self.name = name
        self.document = options.get('document')
        account_id = boto3.client('sts').get_caller_identity().get('Account')
        self.arn = f"arn:aws:iam::{account_id}:policy/{name}"
        super().__init__(**options)
        self.iam = boto3.client('iam')

    def __str__(self):
        return f"IAM policy {self.name}"

    def is_setup(self):
        try:
            self.iam.get_policy(PolicyArn=self.arn)
            return True
        except ClientError:
            return False

    def set_it_up(self):
        self.iam.create_policy(PolicyName=self.name, PolicyDocument=self.document)

    def tear_it_down(self):
        for version in self.iam.list_policy_versions(PolicyArn=self.arn)['Versions']:
            if not version['IsDefaultVersion']:
                self.iam.delete_policy_version(PolicyArn=self.arn, VersionId=version['VersionId'])
        self.iam.delete_policy(PolicyArn=self.arn)


class InstanceProfile(Component):

    def __init__(self, name, **options):
        self.name = name
        super().__init__(**options)
        self.iam = boto3.client('iam')

    def __str__(self):
        return f"IAM Instance Profile {self.name}"

    def is_setup(self):
        try:
            self.iam.get_instance_profile(InstanceProfileName=self.name)
            return True
        except ClientError:
            return False

    def set_it_up(self):
        self.iam.create_instance_profile(InstanceProfileName=self.name)

    def tear_it_down(self):
        self.iam.delete_instance_profile(InstanceProfileName=self.name)


class InstanceProfileRoleAttachment(Component):
    def __init__(self, instance_profile_name, role_name, **options):
        self.instance_profile_name = instance_profile_name
        self.role_name = role_name
        super().__init__(**options)
        self.iam = boto3.client('iam')

    def __str__(self):
        return f"IAM Instance Profile Role Attachment for {self.instance_profile_name}"

    def is_setup(self):
        try:
            response = self.iam.list_instance_profiles_for_role(RoleName=self.role_name)
            instance_profiles = [ip['InstanceProfileName'] for ip in response['InstanceProfiles']]
            return self.instance_profile_name in instance_profiles
        except ClientError:
            return False

    def set_it_up(self):
        self.iam.add_role_to_instance_profile(
            InstanceProfileName=self.instance_profile_name,
            RoleName=self.role_name
        )

    def tear_it_down(self):
        self.iam.remove_role_from_instance_profile(
            InstanceProfileName=self.instance_profile_name,
            RoleName=self.role_name
        )
