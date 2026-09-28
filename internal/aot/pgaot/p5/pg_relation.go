package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_relation_filenode(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_SearchSysCache1(m, int32(57), v7&int64(4294967295))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v10 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
	return int64(0)
L4:
	;
	goto L5
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
	v22 = v20 + v21
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+119)))
	switch v23 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L8
	default:
		goto L9
	}
L6:
	;
	return base.I64_extend_i32_u(v143)
L7:
	;
	v139 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v139)
	return int64(0)
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+88))
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_ReleaseCatCache(m, v10)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	F_ReleaseCatCache(m, v10)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v31 = base.I32_wrap_i64(v7)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+117)))
	v33 = int32(0)
	if v32 == v33 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v143 = v28
	goto L6
L15:
	;
	F_ReleaseCatCache(m, v10)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L53
	}
L16:
	;
	goto L15
L17:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v135 = v130
	goto L16
L18:
	;
	v38 = int32(0)
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filenode[0]))
	if v38 < v40 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v81 = int32(0)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filenode[1]))
	if v81 < v83 {
		goto L39
	} else {
		goto L40
	}
L21:
	;
	v44 = v38
	goto L24
L22:
	;
	goto L23
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filenode[2]))
	if v63 <= int32(0) {
		v135 = v33
		goto L16
	} else {
		goto L30
	}
L24:
	;
	v49 = v44 << (uint(int32(3)) % 32)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_c_F_pg_relation_filenode[3])))
	if v50 == v31 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L23
L26:
	;
	v129 = v49 + int32(_a_F_pg_relation_filenode_0)
	goto L17
L27:
	;
	goto L28
L28:
	;
	v55 = v44 + int32(1)
	if v55 != v40 {
		v44 = v55
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v68 = int32(0)
	goto L31
L31:
	;
	v73 = v68 << (uint(int32(3)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_pg_relation_filenode[4])))
	if v74 != v31 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v129 = v73 + int32(_a_F_pg_relation_filenode_1)
	goto L17
L33:
	;
	v77 = v68 + int32(1)
	if v63 != v77 {
		v68 = v77
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	v135 = v33
	goto L16
L37:
	;
	v111 = int32(0)
	goto L47
L38:
	;
	v129 = v92 + int32(_a_F_pg_relation_filenode_2)
	goto L17
L39:
	;
	v87 = v81
	goto L42
L40:
	;
	goto L41
L41:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filenode[5]))
	if v104 <= int32(0) {
		v135 = v33
		goto L16
	} else {
		goto L46
	}
L42:
	;
	v92 = v87 << (uint(int32(3)) % 32)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+uint32(_c_F_pg_relation_filenode[6])))
	if v31 == v93 {
		goto L38
	} else {
		goto L44
	}
L43:
	;
	goto L41
L44:
	;
	v96 = v87 + int32(1)
	if v96 != v83 {
		v87 = v96
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	goto L37
L47:
	;
	v116 = v111 << (uint(int32(3)) % 32)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+uint32(_c_F_pg_relation_filenode[7])))
	if v117 != v31 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v129 = v116 + int32(_a_F_pg_relation_filenode_3)
	goto L17
L49:
	;
	v120 = v111 + int32(1)
	if v104 != v120 {
		v111 = v120
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	v135 = v33
	goto L16
L53:
	;
	if v135 != 0 {
		v143 = v135
		goto L6
	} else {
		goto L54
	}
L54:
	;
	goto L7
}
