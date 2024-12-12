variable "name" {
  type        = string
  description = "The name of the secret"
}

variable "value" {
  type        = string
  description = "The value of the secret"
}

variable "tags" {
  type        = map(string)
  description = "The tags for the secret"
  default     = {}
}

variable "key_vault_id" {
  type        = string
  description = "Id of the key vault to which secrets need to be added."
}
