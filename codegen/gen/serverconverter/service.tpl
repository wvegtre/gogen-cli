{{define "convert_service"}}

type {{.service_prefix}}Service struct {
	DBOperation   *database.Operation
	DB{{.service_prefix}}Service *{{.group}}.{{.service_prefix}}Service
}

func New{{.service_prefix}}Service() *{{.service_prefix}}Service {
	return &{{.service_prefix}}Service{
		DBOperation:   database.NewOperation(),
		DB{{.service_prefix}}Service: {{.group}}.New{{.service_prefix}}Service(),
	}
}

func (s *{{.service_prefix}}Service) Get(ctx context.Context, id int64) ({{.group}}.{{.model}}, error) {
	var db{{.model}} {{.group}}.{{.model}}
	err := s.DBOperation.QueryByID(ctx, s.DB{{.service_prefix}}Service, id, &db{{.model}})
	if err != nil {
		return db{{.model}}, errors.WithStack(err)
	}
	return db{{.model}}, nil
}

func (s *{{.service_prefix}}Service) List(ctx context.Context, args List{{.service_prefix}}Args) ([]{{.group}}.{{.model}}, error) {
	var db{{.model}}s []{{.group}}.{{.model}}
	whereArgs := args.toDbQueryArgs()
	err := s.DBOperation.Query(ctx, s.DB{{.service_prefix}}Service, whereArgs, &db{{.model}}s)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return db{{.model}}s, nil
}

func (s *{{.service_prefix}}Service) Create(ctx context.Context, args Create{{.service_prefix}}Args) ({{.group}}.{{.model}}, error) {
	var db{{.model}} {{.group}}.{{.model}}
	_, err := s.DBOperation.Create(ctx, s.DB{{.service_prefix}}Service, &db{{.model}})
	if err != nil {
		return db{{.model}}, errors.WithStack(err)
	}
	return db{{.model}}, nil
}

func (s *{{.service_prefix}}Service) UpdateByID(ctx context.Context, id int64, args Update{{.service_prefix}}Args) ({{.group}}.{{.model}}, error) {
	var db{{.model}} {{.group}}.{{.model}}
	_, err := s.DBOperation.UpdateByID(ctx, s.DB{{.service_prefix}}Service, id, args.toDbUpdateArgs(), &db{{.model}})
	if err != nil {
		return db{{.model}}, errors.WithStack(err)
	}
	return db{{.model}}, nil
}

func (s *{{.service_prefix}}Service) Delete(ctx context.Context, id int64) ({{.group}}.{{.model}}, error) {
	var db{{.model}} {{.group}}.{{.model}}
	_, err := s.DBOperation.DeleteByID(ctx, s.DB{{.service_prefix}}Service, id, &db{{.model}})
	if err != nil {
		return db{{.model}}, errors.WithStack(err)
	}
	return db{{.model}}, nil
}

{{- end}}