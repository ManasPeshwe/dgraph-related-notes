# Migrating between two clusters 

With a wide range of use cases that dgraph's database can help with, One might need to migrate the data from one cluster to another. Now, the clusters can be of numerous types. ACL enabled, Non ACL, Dedicated , shared etc. 
In this post we shall discuss the steps required to migrate the data from cluster of type A to type B.

# ACL to Non-ACL  
Dgraph ACLs is an enterprised feature that can be enabled by using the `--acl`  superflag. The use of this feature creates certain predicates and types in the database. 
To demonstrate this, We will spin up an ACL enabled cluster and check the schema and data files. 

### Step 1 : Spin up an ACL enabled cluster
- Create the secret file 
	1. Generate a data encryption key that is 32 bytes long: 
		`tr -dc 'a-zA-Z0-9' < /dev/urandom | dd bs=1 count=32 of=enc_key_file`
	2.  To view the secret key value use `cat enc_key_file`
	3. `echo '<SECRET KEY VALUE>' > hmac_secret_file`
- Start the zero
	-  `dgraph zero --my=localhost:5080`
- Start the alpha
	-  ```dgraph alpha --my=localhost:7080 --zero=localhost:5080 --acl secret-file="/path/to/secret" --security whitelist="<permitted-ip-addresses>"```
### Step 2 : Create an Export
> Note: We have not added any data to the cluster yet. 
- Send login request 
```
curl --location 'localhost:8080/admin' --header 'Content-Type: application/json' --data '{"query":"mutation {\n  login(userId: \"groot\", password: \"password\") {\n    response {\n      accessJWT\n      refreshJWT\n    }\n  }\n}\n","variables":{}}' | jq		
```     
- Send an export request 
```
curl --location --request POST 'localhost:8080/admin' --header 'X-Dgraph-AccessToken: <accessJWT token>' --header 'Content-Type: application/json' --data-raw '{"query":"mutation {\n export(input: {format: \"rdf\"}) {\n\t\tresponse {\n\t\t\tcode\n\t\t\tmessage\n\t\t}\n\t}\n}","variables":{}}'
```
### Step 3: Check the schema and data file 
Now that we have genereated an export, we can find the exported schema and data files in the `/export` directory. 
#### Schema (.schema): 
```
[0x0] <dgraph.xid>:string @index(exact) @upsert . 
[0x0] <dgraph.type>:[string] @index(exact) . 
[0x0] <dgraph.drop.op>:string . 
[0x0] <dgraph.acl.rule>:[uid] . 
[0x0] <dgraph.password>:password . 
[0x0] <dgraph.user.group>:[uid] @reverse . 
[0x0] <dgraph.graphql.xid>:string @index(exact) @upsert . 
[0x0] <dgraph.graphql.schema>:string . 
[0x0] <dgraph.rule.predicate>:string @index(exact) @upsert . 
[0x0] <dgraph.graphql.p_query>:string @index(sha256) . 
[0x0] <dgraph.rule.permission>:int . 
[0x0] type <dgraph.graphql> {
	dgraph.graphql.schema
	dgraph.graphql.xid
}
[0x0] type <dgraph.type.Rule> {
	dgraph.rule.predicate
	dgraph.rule.permission
}
[0x0] type <dgraph.type.User> {
	dgraph.xid
	dgraph.password
	dgraph.user.group
}
[0x0] type <dgraph.type.Group> {
	dgraph.xid
	dgraph.acl.rule
}
[0x0] type <dgraph.graphql.persisted_query> {
	dgraph.graphql.p_query
}
```
#### Data (.rdf) :
```
<0x1> <dgraph.xid> "guardians"^^<xs:string> <0x0> .
<0x2> <dgraph.xid> "groot"^^<xs:string> <0x0> .
<0x1> <dgraph.type> "dgraph.type.Group"^^<xs:string> <0x0> .
<0x2> <dgraph.type> "dgraph.type.User"^^<xs:string> <0x0> .
<0x2> <dgraph.password> "<password-string-hash>"^^<xs:password> <0x0> .
<0x2> <dgraph.user.group> <0x1> <0x0> .
``` 
> Note:  The predicates mentioned in some tripples in the .rdf file like **`dgraph.password`** and **`dgraph.user.group`** are ACL related predicates which are created by the system.

To import the data without any errors we should do two things
1. Remove the **`dgraph.*`** predicates and type defintions except the **`dgraph.type`** and `dgraph.xid` from the schema file.
2. Remove the tripples containing the ACL related predicates. 

Post this **`dgraph live -s "/path/to/schema-file.schema -f "/path/to/data-file.rdf"`** should work without any errors. 

