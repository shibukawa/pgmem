package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_is_standard_join_alias_expression(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	v3 = int32(0)
	if l0 == v3 {
		v146 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v146
L2:
	;
	v8 = l0
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	switch v13 - int32(6) {
	case 0:
		goto L9
	case 1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 24, 25, 26, 27, 28, 29, 30, 31:
		v146 = v3
		goto L1
	case 9:
		goto L8
	case 21, 22, 23:
		goto L7
	case 32:
		goto L6
	default:
		goto L10
	}
L4:
	;
	v146 = v3
	goto L1
L5:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	if v141 != 0 {
		v8 = v141
		goto L3
	} else {
		goto L55
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v37 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	v140 = v8 + int32(4)
	goto L5
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v28 != int32(2) {
		v146 = v3
		goto L1
	} else {
		goto L14
	}
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v23 != v24 {
		v146 = v3
		goto L1
	} else {
		goto L13
	}
L10:
	;
	if v13 != int32(319) {
		v146 = v3
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v18 != v19 {
		v146 = v3
		goto L1
	} else {
		goto L12
	}
L12:
	;
	return int32(1)
L13:
	;
	return int32(1)
L14:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	if v31 == int32(0) {
		v146 = v3
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v140 = v34
	goto L5
L16:
	;
	return int32(1)
L17:
	;
	goto L18
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v43 <= int32(0) {
		v146 = int32(1)
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v46 = int32(0)
	if v46 < v43 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v50 = v43
	goto L22
L21:
	;
	v50 = v46
	goto L22
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v52 = v46
	goto L23
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v51+v52<<(uint(int32(2))%32))))
	v61 = int32(0)
	if v60 == v61 {
		v128 = v61
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v146 = v134
	goto L1
L25:
	;
	if v134 == int32(0) {
		v146 = v134
		goto L1
	} else {
		goto L53
	}
L26:
	;
	v134 = v128
	goto L25
L27:
	;
	v66 = v60
	goto L28
L28:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	switch v71 - int32(6) {
	case 0:
		goto L34
	case 1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 24, 25, 26, 27, 28, 29, 30, 31:
		v128 = v61
		goto L26
	case 9:
		goto L33
	case 21, 22, 23:
		goto L32
	case 32:
		goto L31
	default:
		goto L35
	}
L29:
	;
	v128 = v61
	goto L26
L30:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v123 != 0 {
		v66 = v123
		goto L28
	} else {
		goto L52
	}
L31:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	if v93 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L32:
	;
	v122 = v66 + int32(4)
	goto L30
L33:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	if v84 != int32(2) {
		v128 = v61
		goto L26
	} else {
		goto L39
	}
L34:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v66)+28))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v80 != v81 {
		v128 = v61
		goto L26
	} else {
		goto L38
	}
L35:
	;
	if v71 != int32(319) {
		v128 = v61
		goto L26
	} else {
		goto L36
	}
L36:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v76 != v77 {
		v128 = v61
		goto L26
	} else {
		goto L37
	}
L37:
	;
	v134 = int32(1)
	goto L25
L38:
	;
	v134 = int32(1)
	goto L25
L39:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v66)+28))
	if v87 == int32(0) {
		v128 = v61
		goto L26
	} else {
		goto L40
	}
L40:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v122 = v90
	goto L30
L41:
	;
	v134 = int32(1)
	goto L25
L42:
	;
	goto L43
L43:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v98 <= int32(0) {
		v128 = int32(1)
		goto L26
	} else {
		goto L44
	}
L44:
	;
	v101 = int32(0)
	if v101 < v98 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v105 = v98
	goto L47
L46:
	;
	v105 = v101
	goto L47
L47:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v107 = v101
	goto L48
L48:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v106+v107<<(uint(int32(2))%32))))
	v116 = F_is_standard_join_alias_expression(m, v115, l1)
	mBase = m.M
	if v116 == int32(0) {
		v128 = v116
		goto L26
	} else {
		goto L50
	}
L49:
	;
	v128 = v116
	goto L26
L50:
	;
	v120 = v107 + int32(1)
	if v120 != v105 {
		v107 = v120
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	goto L29
L53:
	;
	v138 = v52 + int32(1)
	if v138 != v50 {
		v52 = v138
		goto L23
	} else {
		goto L54
	}
L54:
	;
	goto L24
L55:
	;
	goto L4
}
