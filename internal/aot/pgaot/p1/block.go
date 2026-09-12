package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BlockRefTableReaderGetBlocks(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	v10 = l0 + int32(65580)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[791])))
	v16 = int32(0)
	v17 = v11
	goto L1
L1:
	;
	if v17 == int32(0) {
		v105 = v16
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v105
L3:
	;
	if base.Ui32(l2) <= base.Ui32(v105) {
		goto L23
	} else {
		goto L24
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1247])))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1248])))
	v24 = int32(1)
	v25 = v17 - v24
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+v25<<(uint(v24)%32)))))
	if v29 != int32(4096) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if base.Ui32(v29) <= base.Ui32(v22) {
		v105 = v16
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if base.Ui32(int32(65535)) < base.Ui32(v22) {
		v105 = v16
		goto L3
	} else {
		goto L14
	}
L8:
	;
	if base.Ui32(l2) <= base.Ui32(v16) {
		v105 = v16
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v39 = v22
	v40 = v16
	goto L10
L10:
	;
	v47 = int32(1)
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+v39<<(uint(v47)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(l1+v40<<(uint(int32(2))%32)))) = v25<<(uint(int32(16))%32) | v50
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1247])))
	v55 = v53 + v47
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1247]))) = v55
	v58 = v40 + v47
	if base.Ui32(v29) <= base.Ui32(v55) {
		v105 = v58
		goto L3
	} else {
		goto L12
	}
L11:
	;
	v105 = v58
	goto L3
L12:
	;
	if base.Ui32(v58) < base.Ui32(l2) {
		v39 = v55
		v40 = v58
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	if base.Ui32(l2) <= base.Ui32(v16) {
		v105 = v16
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v69 = v22
	v70 = v16
	goto L16
L16:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+int32(base.Ui32(v69)>>(uint(int32(3))%32))&int32(536870910)))))
	if int32(base.Ui32(v79)>>(uint(v69&int32(15))%32))&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v105 = v93
	goto L3
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+v70<<(uint(int32(2))%32)))) = v69 + v25<<(uint(int32(16))%32)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1247])))
	v93 = v70 + int32(1)
	v94 = v92
	goto L20
L19:
	;
	v93 = v70
	v94 = v69
	goto L20
L20:
	;
	v96 = v94 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1247]))) = v96
	if base.Ui32(int32(65535)) < base.Ui32(v96) {
		v105 = v93
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if base.Ui32(v93) < base.Ui32(l2) {
		v69 = v96
		v70 = v93
		goto L16
	} else {
		goto L22
	}
L22:
	;
	goto L17
L23:
	;
	goto L2
L24:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[791])))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[792])))
	if v110 == v111 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1248])))
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113+v110<<(uint(int32(1))%32)))))
	if v117 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_BlockRefTableRead(m, l0, v10, v117<<(uint(int32(1))%32))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v125 = v110
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1247]))) = int32(0)
	v129 = v125 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[791]))) = v129
	v16 = v105
	v17 = v129
	goto L1
L29:
	;
	return int32(0)
L30:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[791])))
	v125 = v124
	goto L28
}
func F_BlockStateAsString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	if base.Ui32(l0) <= base.Ui32(int32(19)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[142])))
		v11 = v10
	} else {
		v11 = int32(531628)
	}
	return v11
}
func F_CreateBlockRefTableReader(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_palloc0(m, int32(73776))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[1246]))) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[790]))) = int32(433)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[794]))) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(432)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[793]))) = int32(-1)
		F_BlockRefTableRead(m, v10, v7+int32(12), int32(4))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			if v29 != int32(1697321851) {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(1697321851)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_ReportWalSummaryError(m, int32(0), int32(52250), v7)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(16)
					return v10
				}
			} else {
				m.G0 = v7 + int32(16)
				return v10
			}
		}
	}
}
func F_block_sampling_read_stream_next(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v59 float64
	_ = v59
	var v62 int32
	_ = v62
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 float64
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 float64
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v119 int32
	_ = v119
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v9) < base.Ui32(v10) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v16 = base.B2i32(v12 < v13)
	} else {
		v16 = int32(0)
	}
	if v16 != 0 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v19 = v17 - v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v22 = v20 - v21
		if base.Ui32(v19) < base.Ui32(v22) {
			v25 = l1 + int32(16)
			for {
				v36 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
				v38 = v36 ^ v37
				*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = base.I64_rotl(v38, int64(37))
				*(*int64)(unsafe.Add(mBase, uint32(v25))) = v38<<(uint(int64(16))%64) ^ base.I64_rotl(v36, int64(24)) ^ v38
				v59 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v36*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
				mBase = m.M
				if base.F64_eq(v59, float64(0)) != 0 {
					continue
				} else {
					break
				}
				break
			}
			v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v64 = base.F64_convert_i32_s(v19)
			v67 = base.F64_sub(float64(1), base.F64_div(v64, base.F64_convert_i32_u(v22)))
			if base.F64_gt(v67, v59) != 0 {
				v69 = v62
				v71 = v22
				v75 = v67
				for {
					v77 = int32(1)
					v78 = v69 + v77
					v81 = v71 - v77
					v85 = base.F64_mul(v75, base.F64_sub(float64(1), base.F64_div(v64, base.F64_convert_i32_u(v81))))
					if base.F64_lt(v59, v85) != 0 {
						v69 = v78
						v71 = v81
						v75 = v85
						continue
					} else {
						break
					}
					break
				}
				v87 = v78
			} else {
				v87 = v62
			}
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v96 = v87
			v99 = v95
		} else {
			v96 = v21
			v99 = v18
		}
		v104 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v96 + v104
		*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v99 + v104
		v119 = v96
	} else {
		v119 = int32(-1)
	}
	return v119
}
