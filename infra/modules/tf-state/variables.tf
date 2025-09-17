variable "bucket_name" {
    description = "S3 Bucket name"
    type = string
    validation {
        condition = can(regex("^[a-z0-9{1}[a-z0-9.-]{1,61}[a-z0-9]{1}$", var.bucket_name))
        error_message = "The bucket name must not be empty and must follow S3 naming rules."
    }
}