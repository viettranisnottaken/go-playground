package test_runner

//type TestCase[TInput any, TExpect any] struct {
//	input  TInput
//	expect TExpect
//}
//
//func TestRunner[TInput any, TExpect any, TLambda any](testCases []TestCase[TInput, TExpect], lambda TLambda) {
//	for i := 0; i < len(testCases); i++ {
//		testCase := testCases[i]
//		var result bool
//
//		result = lambda(testCase.input)
//
//		if result != testCase["expect"] {
//			fmt.Printf("Test %d failed, returning %t instead of %t\n", i+1, result, testCase["expect"])
//		} else {
//			fmt.Printf("Test %d passed\n", i+1)
//		}
//	}
//}
