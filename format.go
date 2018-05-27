package whatday

import (
	"fmt"
)

// String today article by default format
func (a *Article) String() string {
	return fmt.Sprintf("%s\n%s", a.Title, a.Text)
}
