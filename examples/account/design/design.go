package design

import . "goa.design/goa.v2/design/rest"
import . "goa.design/goa.v2/dsl/rest"

var _ = API("basic", func() {
	Title("Basic REST example")
	Description("A simple example to  help build v2")
	Server("http://localhost:8080")
})

var _ = Service("account", func() {
	Description("Manage accounts")
	Endpoint("create", func() {
		Description("Create new account")
		Payload(CreateAccount)
		Result(Account)
		Error("name_already_taken", NameAlreadyTaken, "Error returned when name is already taken")
		HTTP(func() {
			POST("/orgs/{org_id}")
			Response(StatusCreated, func() {
				Header("Href:Location")
				Body(Account)
			})
			Response(StatusAccepted, func() {
				Header("Href:Location")
				Body(Empty)
			})
			Response("name_already_taken", StatusConflict)
		})
	})
	Endpoint("list", func() {
		Description("List all accounts")
		Payload(ListFilter)
		Result(ArrayOf(Account))
		HTTP(func() {
			GET("/")
			Params(func() {
				Param("filter")
			})
		})
	})
	Endpoint("show", func() {
		Description("Show account by ID")
		Payload(func() {
			Attribute("id", String, "ID of account to show")
		})
		Result(Account)
		HTTP(func() {
			GET("/{id}")
		})
	})
	Endpoint("delete", func() {
		Description("Delete account by IF")
		Payload(func() {
			Attribute("id", String, "ID of account to delete")
		})
		Result(Empty)
		HTTP(func() {
			DELETE("/{id}")
		})
	})
})

var CreateAccount = Type("CreateAccount", func() {
	Description("CreateAccount is the account creation payload")
	Attribute("org_id", Int, "ID of organization that owns newly created account")
	Attribute("name", String, "Name of new account")
	Required("org_id", "name")
})

var Account = Type("Account", func() {
	Description("Account type")
	Reference(CreateAccount)
	Attribute("href", String, "Href to account")
	Attribute("id", String, "ID of account")
	Attribute("org_id")
	Attribute("name")
	Required("href", "id", "org_id", "name")
})

var NameAlreadyTaken = Type("NameAlreadyTaken", func() {
	Description("NameAlreadyTaken is the type returned when creating an account fails because its name is already taken")
	Attribute("message", String, "Message of error")
	Required("message")
})

var ListFilter = Type("ListFilter", func() {
	Description("ListFilter defines an optional list filter")
	Attribute("filter", String, "Filter is the account name prefix filter", func() {
		Example("go")
	})
})