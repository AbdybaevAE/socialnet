package locale_test

import (
	"context"
	"github.com/abdybaevae/socialnet/pkg/lib/pgtestcontainer"
	"github.com/abdybaevae/socialnet/pkg/repos"
	"github.com/abdybaevae/socialnet/pkg/services/admin/lang"
	"github.com/abdybaevae/socialnet/pkg/services/admin/locale"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type LocaleTestSuite struct {
	pgtestcontainer.PostgresqlSuite
	localeService locale.Service
	ctx           context.Context
}

func (s *LocaleTestSuite) SetupTest() {
	s.PostgresqlSuite.SetupTest()
	s.ctx = context.Background()
	s.localeService = locale.NewService(lang.NewService(repos.NewLangRepo(s.DB)), repos.NewLocaleRepo(s.DB))
}

func (s *LocaleTestSuite) TestShould_successfully_add_locale() {
	// given
	args := locale.AddLocale{
		Translations: map[string]string{
			"english_uk":  "cab",
			"english_usa": "taxi",
		},
	}

	// when
	localeId, err := s.localeService.AddLocale(s.ctx, args)

	// then
	s.Require().NoError(err)
	assert.Positive(s.T(), localeId)

	foundEntity, err := s.localeService.GetLocale(s.ctx, localeId)
	s.Require().NoError(err)
	assert.Equal(s.T(), args.Translations, foundEntity.Translations)
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, &LocaleTestSuite{})
}
