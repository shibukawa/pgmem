package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pointerhash_iterate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v10 = v7
	goto L1
L1:
	;
	if v10&int32(1) != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v27
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v23 = v19 & (v20 - int32(1))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v23
	v27 = v18 + v20<<(uint(int32(3))%32)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v29 = v19 & v28
	if v29 == v23 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v31 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v31)
	goto L8
L7:
	;
	goto L8
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+4)))
	if v34 != int32(1) {
		v10 = base.B2i32(v23 == v29)
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L2
}
func F_pointerhash_lookup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = int32(16)
	v12 = (int32(base.Ui32(l1)>>(uint(v8)%32)) ^ l1) * int32(-2048144789)
	v17 = (int32(base.Ui32(v12)>>(uint(int32(13))%32)) ^ v12) * int32(-1028477387)
	v21 = v7 & (int32(base.Ui32(v17)>>(uint(v8)%32)) ^ v17)
	v24 = v6 + v21<<(uint(int32(3))%32)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)))
	if v25 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v30 = v24
	v32 = v21
	goto L5
L4:
	;
	return v30
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v35 == l1 {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v39 = (v32 + int32(1)) & v7
	v42 = v6 + v39<<(uint(int32(3))%32)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+4)))
	if v43 != 0 {
		v30 = v42
		v32 = v39
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
func F_pointerhash_stat(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 float64
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 float64
	_ = v143
	var v147 int64
	_ = v147
	var v151 float64
	_ = v151
	var v152 float64
	_ = v152
	var v153 float64
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v2 = int32(0)
	v12 = float64(0)
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = F_palloc0(m, v19<<(uint(int32(2))%32))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return
	} else {
		v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		if v24 == int64(0) {
			v124 = v2
			v126 = v2
			v127 = v2
			v128 = v2
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v31 = v2
			v34 = v2
			v35 = v2
			for {
				v45 = v28 + v31<<(uint(int32(3))%32)
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+4)))
				if v46 == int32(1) {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
					v51 = int32(16)
					v55 = (int32(base.Ui32(v50)>>(uint(v51)%32)) ^ v50) * int32(-2048144789)
					v60 = (int32(base.Ui32(v55)>>(uint(int32(13))%32)) ^ v55) * int32(-1028477387)
					v64 = v49 & (int32(base.Ui32(v60)>>(uint(v51)%32)) ^ v60)
					v67 = v22 + v64<<(uint(int32(2))%32)
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
					*(*int32)(unsafe.Add(mBase, uint32(v67))) = v68 + int32(1)
					if base.Ui32(v31) < base.Ui32(v64) {
						v75 = base.I32_wrap_i64(v24)
					} else {
						v75 = int32(0)
					}
					v76 = v31 - v64 + v75
					if base.Ui32(v35) < base.Ui32(v76) {
						v78 = v76
					} else {
						v78 = v35
					}
					v81 = v76 + v34
					v82 = v78
				} else {
					v81 = v34
					v82 = v35
				}
				v85 = v31 + int32(1)
				if base.Ui64(base.I64_extend_i32_u(v85)) < base.Ui64(v24) {
					v31 = v85
					v34 = v81
					v35 = v82
					continue
				} else {
					break
				}
				break
			}
			v88 = int32(0)
			v93 = v88
			v95 = v88
			v98 = v88
			for {
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v22+v98<<(uint(int32(2))%32))))
				v110 = v108 - int32(1)
				if base.Ui32(v93) < base.Ui32(v110) {
					v112 = v110
				} else {
					v112 = v93
				}
				if v108 != 0 {
					v113 = v112
				} else {
					v113 = v93
				}
				if base.Ui32(v110) <= base.Ui32(v108) {
					v116 = v110
				} else {
					v116 = int32(0)
				}
				v117 = v116 + v95
				v119 = v98 + int32(1)
				if base.Ui64(base.I64_extend_i32_u(v119)) < base.Ui64(v24) {
					v93 = v113
					v95 = v117
					v98 = v119
					continue
				} else {
					break
				}
				break
			}
			v124 = v113
			v126 = v117
			v127 = v81
			v128 = v82
		}
		F_pfree(m, v22)
		mBase = m.M
		v137 = m.ExcPending
		if v137 != 0 {
			return
		} else {
			v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v138 == int32(0) {
				v151 = v12
				v152 = v12
				v153 = float64(0)
			} else {
				v143 = base.F64_convert_i32_u(v138)
				v147 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				v151 = base.F64_div(base.F64_convert_i32_u(v126), v143)
				v152 = base.F64_div(base.F64_convert_i32_u(v127), v143)
				v153 = base.F64_div(v143, base.F64_convert_i64_u(v147))
			}
			v156 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v157 = m.ExcPending
			if v157 != 0 {
				return
			} else {
				if v156 != 0 {
					v158 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
					v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*float64)(unsafe.Add(mBase, uint32(v17)+48)) = v151
					*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v124
					*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v126
					*(*float64)(unsafe.Add(mBase, uint32(v17)+32)) = v152
					*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v128
					*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v127
					*(*float64)(unsafe.Add(mBase, uint32(v17)+16)) = v153
					*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v159
					*(*int64)(unsafe.Add(mBase, uint32(v17))) = v158
					F_errmsg_internal(m, int32(_a_F_pointerhash_stat_0), v17)
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_pointerhash_stat_1), int32(1144), int32(_a_F_pointerhash_stat_2))
						mBase = m.M
						v176 = m.ExcPending
						if v176 != 0 {
							return
						} else {
							m.G0 = v17 - int32(-64)
							return
						}
					}
				} else {
					m.G0 = v17 - int32(-64)
					return
				}
			}
		}
	}
}
