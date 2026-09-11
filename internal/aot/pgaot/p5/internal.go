package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_internal_yylex(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+72))
	if int32(0) < v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = v6 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+72)) = v10
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v5+v10<<(uint(int32(2))%32))+76))
	v18 = v6*int32(24) + v5
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)+84))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v19
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v18)+76))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v21
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v18)+68))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v23
	return v15
L2:
	;
	goto L3
L3:
	;
	v28 = F_core_yylex(m, l0, l0+int32(16), l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v35 = v33 + v34
	if v35&int32(3) == int32(0) {
		v59 = v35
		goto L8
	} else {
		goto L9
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v92
	switch v28 - int32(265) {
	case 0:
		goto L25
	default:
		v125 = v28
		goto L23
	case 2:
		goto L24
	}
L7:
	;
	v92 = v84 - v35
	goto L6
L8:
	;
	v63 = v59
	goto L17
L9:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v43 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v92 = int32(0)
	goto L6
L11:
	;
	goto L12
L12:
	;
	v48 = v35
	goto L13
L13:
	;
	v52 = v48 + int32(1)
	if v52&int32(3) == int32(0) {
		v59 = v52
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v84 = v52
	goto L7
L15:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v57 != 0 {
		v48 = v52
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v72 = int32(-2139062144)
	if (int32(16843008)-v69|v69)&v72 == v72 {
		v63 = v63 + int32(4)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v78 = v63
	goto L20
L19:
	;
	goto L18
L20:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v82 != 0 {
		v78 = v78 + int32(1)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v84 = v78
	goto L7
L22:
	;
	goto L21
L23:
	;
	return v125
L24:
	;
	v120 = F_pstrdup(m, v35)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L38
	}
L25:
	;
	v96 = int32(265)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	switch v98 - int32(35) {
	case 0:
		goto L26
	default:
		v125 = v96
		goto L23
	case 25:
		goto L28
	case 27:
		goto L27
	}
L26:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v117 != 0 {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v107 != int32(62) {
		v125 = v96
		goto L23
	} else {
		goto L31
	}
L28:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v101 != int32(60) {
		v125 = v96
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+2)))
	if v104 != 0 {
		v125 = v96
		goto L23
	} else {
		goto L30
	}
L30:
	;
	return int32(278)
L31:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+2)))
	if v112 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v113 = int32(265)
	goto L34
L33:
	;
	v113 = int32(279)
	goto L34
L34:
	;
	return v113
L35:
	;
	v118 = int32(265)
	goto L37
L36:
	;
	v118 = int32(35)
	goto L37
L37:
	;
	return v118
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v120
	v125 = int32(267)
	goto L23
}
