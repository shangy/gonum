// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stat_test

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

func ExamplePC() {
	// iris is a truncated sample of the Fisher's Iris dataset.
	n := 10
	d := 4
	iris := mat.NewDense(n, d, []float64{
		5.1, 3.5, 1.4, 0.2,
		4.9, 3.0, 1.4, 0.2,
		4.7, 3.2, 1.3, 0.2,
		4.6, 3.1, 1.5, 0.2,
		5.0, 3.6, 1.4, 0.2,
		5.4, 3.9, 1.7, 0.4,
		4.6, 3.4, 1.4, 0.3,
		5.0, 3.4, 1.5, 0.2,
		4.4, 2.9, 1.4, 0.2,
		4.9, 3.1, 1.5, 0.1,
	})

	// Calculate the principal component direction vectors
	// and variances.
	var pc stat.PC
	ok := pc.PrincipalComponents(iris, nil)
	if !ok {
		return
	}
	fmt.Printf("variances = %.4f\n\n", pc.VarsTo(nil))

	// Project the data onto the first 2 principal components.
	k := 2
	var proj mat.Dense
	proj.Mul(iris, pc.VectorsTo(nil).Slice(0, d, 0, k))

	fmt.Printf("proj = %.4f", mat.Formatted(&proj, mat.Prefix("       ")))

	// Output:
	// variances = [0.1666 0.0207 0.0079 0.0019]
	//
	// proj = ⎡-6.1686   1.4659⎤
	//        ⎢-5.6767   1.6459⎥
	//        ⎢-5.6699   1.3642⎥
	//        ⎢-5.5643   1.3816⎥
	//        ⎢-6.1734   1.3309⎥
	//        ⎢-6.7278   1.4021⎥
	//        ⎢-5.7743   1.1498⎥
	//        ⎢-6.0466   1.4714⎥
	//        ⎢-5.2709   1.3570⎥
	//        ⎣-5.7533   1.6207⎦
}

func ExamplePC_PrincipalComponents_mDS() {
	// dists is the distance between major Australian cities.
	// http://rosetta.reltech.org/TC/v15/Mapping/data/dist-Aus.csv
	n := 8
	dists := mat.NewSymDense(n, []float64{
		0, 1328, 1600, 2616, 1161, 653, 2130, 1161,
		1328, 0, 1962, 1289, 2463, 1889, 1991, 2026,
		1600, 1962, 0, 2846, 1788, 1374, 3604, 732,
		2616, 1289, 2846, 0, 3734, 3146, 2652, 3146,
		1161, 2463, 1788, 3734, 0, 598, 3008, 1057,
		653, 1889, 1374, 3146, 598, 0, 2720, 713,
		2130, 1991, 3604, 2652, 3008, 2720, 0, 3288,
		1161, 2026, 732, 3146, 1057, 713, 3288, 0,
	})
	/*
		n := 4
		dists := mat.NewSymDense(n, []float64{
			0, 4.05, 8.25, 5.57,
			4.05, 0, 2.54, 2.69,
			8.25, 2.54, 0, 2.11,
			5.57, 2.69, 2.11, 0,
		})
	*/

	// Calculate the principal component direction vectors
	// and variances.
	var pc stat.PC
	ok := pc.PrincipalComponents(stat.ClassicalScaling(nil, dists), nil)
	if !ok {
		return
	}

	// Project the data onto the first 2 (Euclidean) dimensions.
	k := 2
	coords := pc.VectorsTo(nil).Slice(0, n, 0, k)
	fmt.Printf("coords = %f", mat.Formatted(coords, mat.Prefix("         ")))

	// Output:
	//
}
