package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_sparsevec_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 float32
	_ = v64
	var v67 int32
	_ = v67
	var v74 float32
	_ = v74
	var v77 int32
	_ = v77
	var v79 float32
	_ = v79
	var v81 float32
	_ = v81
	var v87 int32
	_ = v87
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 float32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 float32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v155 int32
	_ = v155
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(16)
	v24 = v8 + v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v26 = int32(2)
	v28 = v24 + v25<<(uint(v26)%32)
	v30 = v3 + v23
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	v34 = v30 + v31<<(uint(v26)%32)
	if v31 < v25 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return base.B2i32(v155 <= int32(0))
L5:
	;
	v36 = v31
	goto L7
L6:
	;
	v36 = v25
	goto L7
L7:
	;
	if int32(0) < v36 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v42 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	if v25 <= v31 {
		goto L32
	} else {
		goto L33
	}
L11:
	;
	v55 = v42 << (uint(int32(2)) % 32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v30+v55)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v24)))
	if v57 < v59 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v64 = *(*float32)(unsafe.Add(mBase, uint32(v55+v34)))
	if base.F32_lt(v64, float32(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	if v59 < v57 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v67 = int32(-1)
	goto L18
L17:
	;
	v67 = int32(1)
	goto L18
L18:
	;
	v155 = v67
	goto L4
L19:
	;
	v74 = *(*float32)(unsafe.Add(mBase, uint32(v28+v42<<(uint(int32(2))%32))))
	if base.F32_lt(v74, float32(0)) != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v79 = *(*float32)(unsafe.Add(mBase, uint32(v55+v34)))
	v81 = *(*float32)(unsafe.Add(mBase, uint32(v55+v28)))
	if base.F32_lt(v79, v81) != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v77 = int32(1)
	goto L24
L23:
	;
	v77 = int32(-1)
	goto L24
L24:
	;
	v155 = v77
	goto L4
L25:
	;
	v155 = int32(-1)
	goto L4
L26:
	;
	goto L27
L27:
	;
	if base.F32_gt(v79, v81) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v155 = int32(1)
	goto L4
L29:
	;
	goto L30
L30:
	;
	v87 = v42 + int32(1)
	if v87 != v36 {
		v42 = v87
		goto L11
	} else {
		goto L31
	}
L31:
	;
	goto L12
L32:
	;
	if v31 <= v25 {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	v106 = v31 << (uint(int32(2)) % 32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v24+v106)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v109 <= v108 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v114 = *(*float32)(unsafe.Add(mBase, uint32(v106+v28)))
	if base.F32_lt(v114, float32(0)) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v117 = int32(1)
	goto L37
L36:
	;
	v117 = int32(-1)
	goto L37
L37:
	;
	v155 = v117
	goto L4
L38:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v136 < v134 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v134 = v120
	goto L38
L40:
	;
	goto L41
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v123 = v36 << (uint(int32(2)) % 32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v30+v123)))
	if v121 <= v125 {
		v134 = v121
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v130 = *(*float32)(unsafe.Add(mBase, uint32(v123+v34)))
	if base.F32_lt(v130, float32(0)) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v133 = int32(-1)
	goto L45
L44:
	;
	v133 = int32(1)
	goto L45
L45:
	;
	v155 = v133
	goto L4
L46:
	;
	v155 = int32(-1)
	goto L4
L47:
	;
	goto L48
L48:
	;
	v155 = base.B2i32(v134 < v136)
	goto L4
}
