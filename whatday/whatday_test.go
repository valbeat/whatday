package whatday

import (
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

func TestWhatDay(t *testing.T) {
	ctrl := gomock.NewController(t)
	date := time.Date(2018, 2, 22, 0, 0, 0, 0, time.Local)
	client := NewMockClient(ctrl)
	client.EXPECT().ListPath(gomock.Any(), date).Return([]string{"yurai_other.php?MD=4&NM=10088"}, nil)
	client.EXPECT().GetArticle(gomock.Any(), "yurai_other.php?MD=4&NM=10088").Return(
		&Article{
			Title: "猫の日",
			Text:  "２月２２日を猫の鳴き声「ニャン・ニャン・ニャン」ともじって決められた日で、猫の日制定委員会が１９８７年（昭和６２年）に制定した。",
			Url:   EndPoint + "yurai_other.php?MD=4&NM=10088",
		}, nil)
	articles, err := whatday{client: client}.getArticles(date)
	if err != nil {
		t.Error(err)
	}

	expected := Article{
		Title: "猫の日",
		Text:  "２月２２日を猫の鳴き声「ニャン・ニャン・ニャン」ともじって決められた日で、猫の日制定委員会が１９８７年（昭和６２年）に制定した。",
		Url:   EndPoint + "yurai_other.php?MD=4&NM=10088",
	}
	
	actual := articles[0]
	if diff := cmp.Diff(
		actual,
		expected,
	); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
}
