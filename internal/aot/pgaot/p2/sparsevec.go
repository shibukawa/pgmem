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
	var v41 int32
	_ = v41
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v66 float32
	_ = v66
	var v69 int32
	_ = v69
	var v76 float32
	_ = v76
	var v79 int32
	_ = v79
	var v81 float32
	_ = v81
	var v83 float32
	_ = v83
	var v89 int32
	_ = v89
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 float32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 float32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v157 int32
	_ = v157
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
	return base.B2i32(v157 <= int32(0))
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
	v41 = int32(0)
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
	v55 = v41 << (uint(int32(2)) % 32)
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
	v66 = *(*float32)(unsafe.Add(mBase, uint32(v34+v41<<(uint(int32(2))%32))))
	if base.F32_lt(v66, float32(0)) != 0 {
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
	v69 = int32(-1)
	goto L18
L17:
	;
	v69 = int32(1)
	goto L18
L18:
	;
	v157 = v69
	goto L4
L19:
	;
	v76 = *(*float32)(unsafe.Add(mBase, uint32(v28+v41<<(uint(int32(2))%32))))
	if base.F32_lt(v76, float32(0)) != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v81 = *(*float32)(unsafe.Add(mBase, uint32(v55+v34)))
	v83 = *(*float32)(unsafe.Add(mBase, uint32(v55+v28)))
	if base.F32_lt(v81, v83) != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v79 = int32(1)
	goto L24
L23:
	;
	v79 = int32(-1)
	goto L24
L24:
	;
	v157 = v79
	goto L4
L25:
	;
	v157 = int32(-1)
	goto L4
L26:
	;
	goto L27
L27:
	;
	if base.F32_gt(v81, v83) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v157 = int32(1)
	goto L4
L29:
	;
	goto L30
L30:
	;
	v89 = v41 + int32(1)
	if v89 != v36 {
		v41 = v89
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
	v108 = v31 << (uint(int32(2)) % 32)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v24+v108)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v111 <= v110 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v116 = *(*float32)(unsafe.Add(mBase, uint32(v108+v28)))
	if base.F32_lt(v116, float32(0)) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v119 = int32(1)
	goto L37
L36:
	;
	v119 = int32(-1)
	goto L37
L37:
	;
	v157 = v119
	goto L4
L38:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v138 < v137 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v137 = v122
	goto L38
L40:
	;
	goto L41
L41:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v125 = v36 << (uint(int32(2)) % 32)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v30+v125)))
	if v123 <= v127 {
		v137 = v123
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v132 = *(*float32)(unsafe.Add(mBase, uint32(v125+v34)))
	if base.F32_lt(v132, float32(0)) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v135 = int32(-1)
	goto L45
L44:
	;
	v135 = int32(1)
	goto L45
L45:
	;
	v157 = v135
	goto L4
L46:
	;
	v157 = int32(-1)
	goto L4
L47:
	;
	goto L48
L48:
	;
	v157 = base.B2i32(v137 < v138)
	goto L4
}
