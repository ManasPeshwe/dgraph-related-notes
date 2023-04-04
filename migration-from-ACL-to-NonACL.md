# Migration from ACL enabled to Non ACL backend. 

Moving from an ACL enabled to a Non-ACL backend is a 4 step process. 
1. Creation of an export from the source backend (ACL enabled)
2. Remove some internal predicate and type definitions from the schema file. 
3. Remove some triples from the data file which mention the internal predicates
4. Import the data and schema to the target backend (Non-ACL)

Let's discuss the steps in detail  
### Creation of export from the source backend (ACL-enabled) 
**Cloud UI** 
Creating an Export from the cloud UI is straightforward. 
`Settings` -> `Exports` -> `Create Export`

Download the **".schema"** and **".rdf"** files.

**On prem** 
Step 1 :  Make a login call to obtain the **accessJWT** 
```
mutation {
login(userId: "groot", password: "password") 
	{
		response {
			accessJWT, 
			refreshJWT
		}
	}
}
```
Curl request  : 
```
curl --location --request POST 'localhost:8080/admin' --header 'Content-Type: application/json' --data-raw '{"query":"mutation {login(userId: \"groot\", password: \"password\") {response {accessJWT, refreshJWT}}}","variables":{}}'
```
> Note : Copy the accessJWT 

Step 2 : Create an Export 
```
mutation {
	export(input: {format: "rdf"}) 
		{
		response {
			code
			message
		}
	}
}
```
Curl request: 
```
curl --location --request POST '<alpha-host>:8080/admin' --header 'X-Dgraph-AccessToken: <accessJWT-token>' --header 'Content-Type: application/json' --data-raw '{"query":"mutation {export(input: {format: \"rdf\"}) {response {code, message}}}","variables":{}}'
```
> Find the exported files in the /export directory.
#### Remove internal predicate and type definitions from the schema file
Remove the following predicate definitons from the schema file  : 
1. `dgraph.acl.rule`
2. `dgraph.rule.predicate`
3. `dgraph.rule.permission`
4. `dgraph.user.group`
5. `dgraph.user.group`

Remove the following type definitions from the schema file  : 
1. `dgraph.type.Rule`
2. `dgraph.type.User`
3. `dgraph.type.Group `

#### Remove triples from the data file which mention the internal predicates

Remove the tripples which mention the following predicates. 
1. **`dgraph.*`** **EXCEPT** `dgraph.type` and `dgraph.xid`
2. dgraph.type.User , dgraph.type.Rule 

### Import the data and schema to the target backend (Non-ACL) 
To import the data schema you can use  **`dgraph live`** or **`dgraph bulk`**

Documentation links  :
- dgraph live : https://dgraph.io/docs/howto/importdata/live-loader/
- dgraph bulk: https://dgraph.io/docs/howto/importdata/bulk-loader/
