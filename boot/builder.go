package boot

type Builder interface {
	To(to string) Builder
	Name(name string) Builder
	Tags(tags ...string) Builder
	Description(desc string) Builder
}

type chainBuilder interface {
	build(e *Engine) error
}

type _builder struct {
	to       string
	name     string
	desc     string
	tags     []string
	resolver interface{}
}

func (my *_builder) To(to string) Builder            { my.to = to; return my }
func (my *_builder) Name(name string) Builder        { my.name = name; return my }
func (my *_builder) Tags(tags ...string) Builder     { my.tags = tags; return my }
func (my *_builder) Description(desc string) Builder { my.desc = desc; return my }

func (my *_builder) build(e *Engine) error {
	return e.AddTo(my.resolver, my.to, my.name, my.desc)
}

func (my *Engine) NewBuilder(resolver interface{}) Builder {
	b := &_builder{resolver: resolver, name: getFuncName(resolver)}
	my.builders = append(my.builders, b)
	return b
}

func (my *Engine) NewQuery(resolver interface{}) Builder {
	return my.NewBuilder(resolver).To(Query)
}

func (my *Engine) NewMutation(resolver interface{}) Builder {
	return my.NewBuilder(resolver).To(Mutation)
}

func (my *Engine) NewSubscription(resolver interface{}) Builder {
	return my.NewBuilder(resolver).To(Subscription)
}
