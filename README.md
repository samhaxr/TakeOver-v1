# TakeOver-v1

### What is Subdomain Takeover?

Subdomain takeover is a class of vulnerability where subdomain points to an external service that has been deleted. The external services are Github, Heroku, Gitlab, Tumblr and so on. Let’s assume we have a subdomain sub.example.com that points to an external service such as GitHub. If the Github page is removed by its owner and forgot to remove the DNS entry that points to GitHub service. An attacker can simply takeover subdomain by adding CNAME file containing the sub.example.com.

Here is the command that checks CNAME record of a subdomain.

$dig CNAME apt.shopify.com --> apt.shopify.com.s3-website-us-east-1.amazonaws.com.

### How Can Takeover script help bug bounty hunters?

There are a lot of sites having thousands of subdomain and it’s really hard to check each subdomain. Here we got a script that shows CNAME record for each domain. It takes a file name as an input and perform some action and finally produce it output, which shows CNAME record for each domain. The input file should contain a list of subdomains.

### How can I recognise if the subdomain is vulnerable to subdomain takeover?

There are some fingerprints should be analysed when service is deleted and DNS entry remains as it is. The attacker get this error when visiting vulnerable subdomain such as “There isn’t a Github Pages site here.” or view below image for more detail.

![Alt text](/53D457E2-57EC-4217-89EA-9A6DA84E1ACE.png?raw=true "TakeOver-v1" )

Security researcher @edoverflow has listed all services and their fingerprints. For more detail visit https://github.com/EdOverflow/can-i-take-over-xyz

# Youtube

[![Takeover-v1](https://img.youtube.com/vi/vpCI8foxP00/0.jpg)](https://www.youtube.com/watch?v=vpCI8foxP00)
