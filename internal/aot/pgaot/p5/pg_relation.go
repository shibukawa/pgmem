package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_relation_filenode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_SearchSysCache1(m, int32(57), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v138
L2:
	;
	v138 = int32(0)
	goto L1
L3:
	;
	return int32(0)
L4:
	;
	if v8 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v14 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
	goto L2
L6:
	;
	goto L7
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
	v18 = v16 + v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+119)))
	switch v19 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L9
	default:
		goto L10
	}
L8:
	;
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v133)
	goto L2
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	if v24 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_ReleaseCatCache(m, v8)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	F_ReleaseCatCache(m, v8)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+117)))
	v29 = int32(0)
	if v28 == v29 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	return v24
L16:
	;
	F_ReleaseCatCache(m, v8)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L47
	}
L17:
	;
	goto L16
L18:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	v128 = v123
	goto L17
L19:
	;
	v106 = int32(0)
	goto L43
L20:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[1044]))
	if int32(0) < v35 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[1045]))
	if int32(0) < v63 {
		goto L31
	} else {
		goto L32
	}
L23:
	;
	v40 = v29
	goto L26
L24:
	;
	goto L25
L25:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[1046]))
	if v58 <= int32(0) {
		v128 = v29
		goto L17
	} else {
		goto L30
	}
L26:
	;
	v44 = v40 << (uint(int32(3)) % 32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[1047])))
	if v7 == v47 {
		v119 = v44 + int32(4459888)
		goto L18
	} else {
		goto L28
	}
L27:
	;
	goto L25
L28:
	;
	v50 = v40 + int32(1)
	if v50 != v35 {
		v40 = v50
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L19
L31:
	;
	v68 = v29
	goto L34
L32:
	;
	goto L33
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[1048]))
	if v86 <= int32(0) {
		v128 = v29
		goto L17
	} else {
		goto L38
	}
L34:
	;
	v72 = v68 << (uint(int32(3)) % 32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+uint32(_consts[1049])))
	if v7 == v75 {
		v119 = v72 + int32(4458840)
		goto L18
	} else {
		goto L36
	}
L35:
	;
	goto L33
L36:
	;
	v78 = v68 + int32(1)
	if v78 != v63 {
		v68 = v78
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v92 = int32(0)
	goto L39
L39:
	;
	v96 = v92 << (uint(int32(3)) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+uint32(_consts[1050])))
	if v7 == v99 {
		v119 = v96 + int32(4459364)
		goto L18
	} else {
		goto L41
	}
L40:
	;
	v128 = v29
	goto L17
L41:
	;
	v102 = v92 + int32(1)
	if v86 != v102 {
		v92 = v102
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v110 = v106 << (uint(int32(3)) % 32)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)+uint32(_consts[1051])))
	if v7 == v113 {
		v119 = v110 + int32(4460412)
		goto L18
	} else {
		goto L45
	}
L44:
	;
	v128 = v29
	goto L17
L45:
	;
	v116 = v106 + int32(1)
	if v58 != v116 {
		v106 = v116
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	if v128 != 0 {
		v138 = v128
		goto L1
	} else {
		goto L48
	}
L48:
	;
	goto L8
}
