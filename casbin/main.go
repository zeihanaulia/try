package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize a Xorm adapter with Postgre database.
	a, err := xormadapter.NewAdapter("postgres", "user=postgres password=postgrespw dbname=casbin port=49153 sslmode=disable")
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}

	// Setting model basic ACL
	m, err := model.NewModelFromString(`
[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.act
`)
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}

	// Initiate enforcer
	e, err := casbin.NewEnforcer(m, a)
	fmt.Println(e, err)

	// Add new policies to db
	ok, err := e.AddPolicies([][]string{
		// represent policy sales-order-return resource
		{
			"role:sales-order", "oms", "sales-order", "read",
		},
		{
			"role:sales-order", "oms", "sales-order", "write",
		},
		{
			"role:sales-order-approval", "oms", "sales-order", "approval",
		},

		// represent policy sales-order-return resource
		{
			"role:sales-order-return", "oms", "sales-order-return", "read",
		},
		{
			"role:sales-order-return", "oms", "sales-order-return", "write",
		},
		{
			"role:sales-order-return-approval", "oms", "sales-order-return", "approval",
		},
	})
	fmt.Println("Add Policies", ok, err)

	ok, err = e.AddGroupingPolicies([][]string{
		{
			"role:staff-1", "role:sales-order", "oms",
		},
		{
			"role:staff-2", "role:sales-order-return", "oms",
		},
		{
			"role:manager", "role:staff-1", "oms",
		},
		{
			"role:manager", "role:staff-2", "oms",
		},
		{
			"role:manager", "role:sales-order-approval", "oms",
		},
		{
			"role:manager", "role:sales-order-return-approval", "oms",
		},
		{
			"martin", "role:staff-1", "oms",
		},
		{
			"bob", "role:staff-2", "oms",
		},
		{
			"alice", "role:manager", "oms",
		},
	})
	fmt.Println("Add Grouping Policies", ok, err)

	sub := "alice" // the user that wants to access a resource.
	dom := "oms"
	obj := "sales-order" // the resource that is going to be accessed.
	act := "read"        // the operation that the user performs on the resource.

	ok, err = e.Enforce(sub, dom, obj, act)
	fmt.Println("Enforce", ok, err)
}
