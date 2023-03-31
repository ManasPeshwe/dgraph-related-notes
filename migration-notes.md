# Migrating between two clusters 

With a wide range of use cases that dgraph's database can help with, One might need to migrate the data from one cluster to another. Now, the clusters can be of numerous types. ACL enabled, Non ACL, Dedicated , shared etc. 
In this post we shall discuss the steps required to migrate the data from cluster of type A to type B.

# ACL to Non-ACL  
Dgraph ACLs is an enterprised feature that can be enabled by using the `--acl`  superflag. The use of this feature creates certain predicates and types in the database. 

Before moving from an ACL enabled backend to non ACL backend some changes are required to be made in the schema and data files 
	1. Remove predicate and type definitions for `dgraph.acl.*` , `dgraph.type.user`,`dgraph.type.group` from the schema file. 
	2. Remove all tripples with `dgraph.*` **except** `dgraph.graphql.*` , `dgraph.type` ,`dgraph.xid` , `dgraph.type.user` , `dgraph.type.group`

#### Things to remove from the schema file : 
Remove type definitions from schema
```
​​type <dgraph.type.Rule> {
	dgraph.rule.predicate
	dgraph.rule.permission
}
type <dgraph.type.User> {
	dgraph.xid
	dgraph.password
	dgraph.user.group
}
type <dgraph.type.Group> {
	dgraph.xid
	dgraph.acl.rule
}
```
Remove the following predicate definitions :
```
<dgraph.acl.rule>:[uid] . 
<dgraph.password>:password . 
<dgraph.user.group>:[uid] @reverse . 
<dgraph.rule.predicate>:string @index(exact) @upsert . 
<dgraph.rule.permission>:int . 
```
### Things to remove from the data file :
> Note: the \<uids> can be different in your data file.

**`<0x2> <dgraph.password> "-password-hash-"^^<xs:password> <0x0> .`**
**`<0x2> <dgraph.user.group> <0x1> <0x0> .`**

# ACL to ACL

When moving from an ACL enabled cluster to a non ACL cluster, no changes are needed in the data file but only the schema file. 
1. Remove all type definitions for`dgraph.*` types from the schema.
2. Remove predicate definitions for `<dgraph.acl.rule>` ,`<dgraph.password>` ,`<dgraph.user.group>`, `<dgraph.rule.predicate>`, `<dgraph.rule.permission>`. 
