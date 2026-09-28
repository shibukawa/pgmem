package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_create_setop_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 float64, l7 float64) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v76 int32
	_ = v76
	var v79 float64
	_ = v79
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v95 int32
	_ = v95
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 float64
	_ = v120
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	var v160 float64
	_ = v160
	var v166 int32
	_ = v166
	var v169 float64
	_ = v169
	var v171 int32
	_ = v171
	var v175 float64
	_ = v175
	var v176 float64
	_ = v176
	var v179 float64
	_ = v179
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	v15 = F_palloc0(m, int32(104))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(1610612736315)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v23 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)) = uint8(v23)
		*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v22
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
		if v29 != int32(1) {
			v37 = v23
		} else {
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
			if v33 != int32(1) {
				v37 = int32(0)
			} else {
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
				v37 = v36
			}
		}
		v39 = v37 & int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v15)+21)) = uint8(v39)
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v41 + v42
		if l4 == int32(0) {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
			v48 = v47
		} else {
			v48 = int32(0)
		}
		*(*float64)(unsafe.Add(mBase, uint32(v15)+96)) = l6
		*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v48
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		v58 = v56 + v57
		*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v58
		if l4 == int32(0) {
			v62 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
			v63 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
			*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = base.F64_add(v62, v63)
			v66 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
			v67 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
			v70 = *(*float64)(unsafe.Add(mBase, _c_F_create_setop_path[0]))
			v71 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
			v72 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
			if l5 != 0 {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
				v79 = base.F64_convert_i32_s(v76)
			} else {
				v79 = float64(0)
			}
			*(*float64)(unsafe.Add(mBase, uint32(v15)+56)) = base.F64_add(base.F64_mul(v70, l7), base.F64_add(base.F64_mul(base.F64_mul(v70, base.F64_add(v71, v72)), v79), base.F64_add(v66, v67)))
			*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = l7
			return v15
		} else {
			v86 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
			v87 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
			v90 = *(*float64)(unsafe.Add(mBase, _c_F_create_setop_path[0]))
			v91 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
			v92 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
			if l5 != 0 {
				v95 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
				v98 = base.F64_convert_i32_s(v95)
			} else {
				v98 = float64(0)
			}
			v100 = base.F64_add(base.F64_mul(base.F64_mul(v90, base.F64_add(v91, v92)), v98), base.F64_add(v86, v87))
			*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v100
			*(*float64)(unsafe.Add(mBase, uint32(v15)+56)) = base.F64_add(base.F64_mul(v90, l7), v100)
			v106 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_setop_path[1])))
			if v106 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v58 + int32(1)
			} else {
			}
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+32))
			v118 = int32(-1)
			v120 = base.F64_div(l6, float64(0.9))
			if base.F64_ge(v120, float64(4.294967296e+09)) != 0 {
				v166 = v118
			} else {
				v123 = int64(2)
				v124 = base.I64_trunc_sat_f64_u(v120)
				if base.Ui64(v124) <= base.Ui64(v123) {
					v127 = v123
				} else {
					v127 = v124
				}
				v128 = int64(1)
				if v127&(v127-v128) == int64(0) {
					v138 = v127
				} else {
					v138 = v128 << (uint(int64(64)-base.I64_clz(v127)) % 64)
				}
				v140 = v138 * int64(12)
				if base.Ui64(int64(2147483646)) < base.Ui64(v140) {
					v166 = v118
				} else {
					v188 = int32(32)
					v160 = base.F64_add(base.F64_mul(l6, base.F64_convert_i32_u((v113+int32(7))&int32(-8)+v188)), base.F64_convert_i32_u(base.I32_wrap_i64(v140)+v188))
					if base.F64_ge(v160, float64(4.294967295e+09)) != 0 {
						v166 = v118
					} else {
						v166 = base.I32_trunc_sat_f64_u(v160)
					}
				}
			}
			v169 = *(*float64)(unsafe.Add(mBase, _c_F_create_setop_path[2]))
			v171 = *(*int32)(unsafe.Add(mBase, _c_F_create_setop_path[3]))
			v175 = base.F64_mul(base.F64_mul(v169, base.F64_convert_i32_s(v171)), float64(1024))
			v176 = float64(4.294967295e+09)
			if base.F64_lt(v175, v176) != 0 {
				v179 = v175
			} else {
				v179 = v176
			}
			if base.Ui32(base.I32_trunc_sat_f64_u(v179)) <= base.Ui32(v166) {
				v182 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v182 + int32(1)
			} else {
			}
			*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = l7
			return v15
		}
	}
}
func F_setop_column_grouping_eqop(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v84
L2:
	;
	v84 = int32(0)
	goto L1
L3:
	;
	v9 = l0
	goto L4
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v13 != int32(142) {
		goto L2
	} else {
		goto L6
	}
L5:
	;
	goto L2
L6:
	;
	if l1 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v31 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	if v16 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 < l1 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+l1<<(uint(int32(2))%32)-int32(4))))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	return v28
L11:
	;
	if v75 != 0 {
		v84 = v75
		goto L1
	} else {
		goto L24
	}
L12:
	;
	v75 = v69
	goto L11
L13:
	;
	v69 = int32(0)
	goto L12
L14:
	;
	v38 = v31
	goto L15
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v42 != int32(142) {
		goto L13
	} else {
		goto L17
	}
L16:
	;
	goto L13
L17:
	;
	if l1 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v60 = F_setop_column_grouping_eqop(m, v59, l1)
	mBase = m.M
	if v60 != 0 {
		v69 = v60
		goto L12
	} else {
		goto L22
	}
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	if v45 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v48 < l1 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50+l1<<(uint(int32(2))%32)-int32(4))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v75 = v57
	goto L11
L22:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v61 != 0 {
		v38 = v61
		goto L15
	} else {
		goto L23
	}
L23:
	;
	goto L16
L24:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v76 != 0 {
		v9 = v76
		goto L4
	} else {
		goto L25
	}
L25:
	;
	goto L5
}
