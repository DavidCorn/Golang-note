import s "strings"

s.Contains("test", "est")
s.Count("test", "t")
s.HasPrefix("test", "te")
s.HasSuffix("test", "st")
s.Index("test", "e")
s.Join([]string{"a", "b"}, "-") // a-b
s.Repeat("a", 5)
s.Replace("foo", "o", "0", -1) // -1 means replace all "o"
s.Split("a-b-c-d-e", "-")
s.ToLower("TEST")
s.ToUpper("test")
s.TrimSpace("   test   \n")

