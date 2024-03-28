terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~>4.16"
    }
  }
}

provider "aws" {
  shared_credentials_file = "/home/ubuntu/.aws/credentials"
  profile                 = "default"
  region                  = "ap-south-1"
}

resource "aws_s3_bucket" "creation-storage" {
  bucket        = "hyper-hive-data"
  force_destroy = true
  # acl    = "public-read" 

  tags = {
    Name       = "HyperHive"
    enviroment = "development"
  }
}

resource "aws_s3_bucket_ownership_controls" "bucket-owner" {
  bucket = aws_s3_bucket.creation-storage.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_public_access_block" "bucket-access" {
  bucket = aws_s3_bucket.creation-storage.id

  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}

resource "aws_s3_bucket_acl" "bucker-acl" {
  depends_on = [
    aws_s3_bucket_ownership_controls.bucket-owner,
    aws_s3_bucket_public_access_block.bucket-access,
  ]

  bucket = aws_s3_bucket.creation-storage.id
  acl    = "public-read"
}

resource "aws_s3_object" "subfolder-one" {
  bucket = aws_s3_bucket.creation-storage.id
  key    = "profile images/"
}

resource "aws_s3_bucket_public_access_block" "creation-storage" {
  bucket = aws_s3_bucket.creation-storage.id

  block_public_acls   = false
  block_public_policy = false
}