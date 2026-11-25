package charset

import (
	. "github.com/pingcap/check"
)

func (s *testCharsetSuite) TestMySQL8UTF8MB40900Collations(c *C) {
	for _, name := range []string{
		"utf8mb4_0900_as_cs",
		"utf8mb4_0900_as_ci",
		"utf8mb4_0900_bin",
		"utf8mb4_zh_0900_as_cs",
		"utf8mb4_ja_0900_as_cs_ks",
	} {
		col, err := GetCollationByName(name)
		c.Assert(err, IsNil, Commentf("name=%s", name))
		c.Assert(col.Name, Equals, name)
		c.Assert(col.CharsetName, Equals, "utf8mb4")
		c.Assert(col.ID > 255, Equals, true, Commentf("name=%s id=%d", name, col.ID))
	}
	_, err := GetCollationByName("utf8mb4_0900_as_xx")
	c.Assert(err, NotNil)
}
