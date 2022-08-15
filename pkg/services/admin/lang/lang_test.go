package lang_test

import (
	"context"
	"github.com/abdybaevae/socialnet/pkg/entities"
	"github.com/abdybaevae/socialnet/pkg/lib/pgtestcontainer"
	"github.com/abdybaevae/socialnet/pkg/repos"
	"github.com/abdybaevae/socialnet/pkg/services/admin/lang"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"testing"
)

type LangTestSuite struct {
	pgtestcontainer.PostgresqlSuite
	langService lang.Service
	ctx         context.Context
}

func (s *LangTestSuite) SetupTest() {
	s.PostgresqlSuite.SetupTest()
	s.langService = lang.NewService(repos.NewLangRepo(s.DB))
	s.ctx = context.Background()
}

func (s *LangTestSuite) TestShould_add_language_into_database() {
	// given
	args := lang.CreateLangArgs{
		Code: "kazakh",
		Name: "Kazakh language",
	}

	// when
	s.Require().NoError(s.langService.CreateLang(s.ctx, args))

	// then
	var langs []entities.LanguageEntity
	err := s.DB.SelectContext(s.ctx, &langs, "select * from lang")
	s.Require().NoError(err)
	assert.EqualValues(s.T(), len(s.AlreadyExistedLangs)+1, len(langs))

	// then
	s.Require().NoError(s.DB.SelectContext(s.ctx, &langs, "select * from lang where lang_code = $1", args.Code))
	foundLang := langs[0]
	assert.EqualValues(s.T(), args.Name, foundLang.Name)
}
func (s *LangTestSuite) TestShould_retrieve_all_languages_from_database_and_return_exact_values() {
	// given
	givenLangs := s.AlreadyExistedLangs

	// when
	foundLangs, err := s.langService.GetLangList(s.ctx)

	// then
	s.Require().NoError(err)
	assert.Equal(s.T(), len(foundLangs), len(givenLangs))

	for _, existedLang := range givenLangs {
		found := false
		for _, foundLang := range foundLangs {
			if existedLang.Id != foundLang.LangId {
				continue
			}
			found = true
			assert.Equal(s.T(), existedLang.Id, foundLang.LangId)
			assert.Equal(s.T(), existedLang.Code, foundLang.LangCode)
			assert.Equal(s.T(), existedLang.Name, foundLang.LangName)
		}
		assert.True(s.T(), found)
	}

}
func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, &LangTestSuite{})
}
