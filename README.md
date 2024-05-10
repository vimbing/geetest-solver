# Geetest v4 slider captcha golang solver

Based on [this gravilk's repo](https://github.com/gravilk/geetest-v4-slide-documented)


# Example
```go
package  main

import (
	"fmt"
	"github.com/vimbing/geetest-solver"
)  

const CAPTCHA_ID =  "3f765ss62g871421d5723a68371726a676"

func  main() {
	solver := geetest.New(CAPTCHA_ID)
	
	token, err := solver.Solve()

	if err !=  nil {
		panic(err)
	}
	
	fmt.Printf("token: %v\n", token)
}
```
```bash
token: eyJjYXB0Y2hhX2lkIjoiM2Y3NjU4ODI4NzA0ZjFkMjcyM2E2ODM3MTcyNmE2NzYiLCJsb3RfbnVtYmVyIjoiZGNkZWNiZDQxZWM0NGUxY2IyZDUyNzQwMjA1YWQwZTAiLCJwYXNzX3Rva2VuIjoiMmY2NTUyZmQ3MWE5MTZkMzcwNjAxZjM5M2IzOTEwMjlmMTkxODBmMTQwNzQ5ODc5NDhiYzMyNGM2OWRjYThjOCIsImdlbl90aW1lIjoiMTcxNTM0MzQ4MSIsImNhcHRjaGFfb3V0cHV0IjoiY2k3RDF5dEx5VmNDVkhmUkszdHdLc01XajlKOXRSSExkWHhBNnVncGJjQXdReG41Zm04YlRjXzludlpiUlA5T2Y0S1VvRVRhZVJzdnZhdnZ4b1hVV1ljX1hiQzFpekFrbnN2dmJoTDdaam5rRUdQOEwya1RJM181YWZpU01vUlZ3TWp3RWtnSEVJRnZsQ1hTdFFMQmtTZkFJeU4yVjNXandSR0JOMkxUdGFTZHdIbU5KblpMaWM4MFNILVpia2lGSG8tUTVFR2w5MlpUVlpNUnptYUJpNVQ5MnM1SkNod3ZGUExGWnNNbFNMYjEtbVdHR3RUWG1HN0hmZy03QlgtUXdwUnlYYW1YZDFCcFk3NFJNWlVxTDQ0OUczZHFfY0JTMF82cF9RZkk3bTZiQXpZTjlRM052Z29sTHBTZHE3T291REpNcnlEWWk5SjBZUTR5Tm9iVlluQ3g3TlVybmVOUnl3eWlEQlFUS01uMVVaaXZ1R012a3VhOWIxT3UwTEdIOGVmeXRQY3Q5TUNYYmNtTjB2czJJZnYxQzdIeGJOYXdaakkwV1RqRmNDbjB6Sk80VDNYLVU0OWlYQWpsSjFoSSJ9
```

## Todo

```
- [ ] Add more image indexes, or change to dynamic slide solver  
```