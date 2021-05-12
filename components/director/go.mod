module github.com/kyma-incubator/compass/components/director

go 1.15

require (
	github.com/99designs/gqlgen v0.11.0
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/dlmiddlecote/sqlstats v1.0.2
	github.com/form3tech-oss/jwt-go v3.2.2+incompatible
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/go-openapi/runtime v0.19.26
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-multierror v1.1.0
	github.com/jmoiron/sqlx v1.3.1
	github.com/kyma-incubator/compass/components/connector v0.0.0-20210315172259-e186b4cac80b
	github.com/kyma-incubator/compass/components/operations-controller v0.0.0-20210301144857-4b0b2ea4c892
	github.com/lestrrat-go/iter v1.0.0
	github.com/lestrrat-go/jwx v1.1.4
	github.com/lib/pq v1.9.0
	github.com/machinebox/graphql v0.2.3-0.20181106130121-3a9253180225
	github.com/moby/moby v20.10.6+incompatible // indirect
	github.com/oliveagle/jsonpath v0.0.0-20180606110733-2e52cf6e6852
	github.com/onrik/logrus v0.8.0
	github.com/ory/hydra-client-go v1.9.2
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.9.0
	github.com/sirupsen/logrus v1.8.0
	github.com/stretchr/testify v1.7.0
	github.com/tidwall/gjson v1.6.8
	github.com/tidwall/sjson v1.1.5
	github.com/vektah/gqlparser/v2 v2.0.1
	github.com/vrischmann/envconfig v1.3.0
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/oauth2 v0.0.0-20210220000619-9bb904979d93
	golang.org/x/sys v0.0.0-20210331175145-43e1dd70ce54 // indirect
	golang.org/x/tools v0.1.0 // indirect
	k8s.io/api v0.17.2 // DO NOT BUMP
	k8s.io/apimachinery v0.17.2 // DO NOT BUMP
	k8s.io/apiserver v0.17.2
	k8s.io/client-go v0.17.2 // DO NOT BUMP
	sigs.k8s.io/controller-runtime v0.5.0 // DO NOT BUMP
)

replace github.com/kyma-incubator/compass/components/operations-controller => ../operations-controller
