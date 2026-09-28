package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bpchar_pattern_ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = int32(1)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v21 = v19 & v17
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v129 != v6 {
		goto L53
	} else {
		goto L54
	}
L5:
	;
	v22 = v17
	goto L7
L6:
	;
	v22 = int32(4)
	goto L7
L7:
	;
	v23 = v22 + v6
	if v19 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v56 = v50
	goto L19
L9:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	if v29 == int32(18) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v40 = int32(1)
	if v21 != 0 {
		v50 = int32(base.Ui32(v19)>>(uint(v40)%32)) - v40
		goto L8
	} else {
		goto L18
	}
L12:
	;
	v32 = int32(16)
	goto L14
L13:
	;
	v32 = int32(0)
	goto L14
L14:
	;
	if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v39 = int32(4)
	goto L17
L16:
	;
	v39 = v32
	goto L17
L17:
	;
	v50 = v39
	goto L8
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L19:
	;
	if v56 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v70 = int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v74 = v72 & v70
	if v74 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	goto L20
L22:
	;
	v68 = v50 & (v50 >> (uint(int32(31)) % 32))
	goto L21
L23:
	;
	goto L24
L24:
	;
	v63 = v56 - int32(1)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v63))))
	if v65 == int32(32) {
		v56 = v63
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v68 = v56
	goto L21
L26:
	;
	v75 = v70
	goto L28
L27:
	;
	v75 = int32(4)
	goto L28
L28:
	;
	v76 = v75 + v11
	if v72 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v109 = v103
	goto L40
L30:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v82 == int32(18) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v93 = int32(1)
	if v74 != 0 {
		v103 = int32(base.Ui32(v72)>>(uint(v93)%32)) - v93
		goto L29
	} else {
		goto L39
	}
L33:
	;
	v85 = int32(16)
	goto L35
L34:
	;
	v85 = int32(0)
	goto L35
L35:
	;
	if base.Ui32((v82-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v92 = int32(4)
	goto L38
L37:
	;
	v92 = v85
	goto L38
L38:
	;
	v103 = v92
	goto L29
L39:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v103 = int32(base.Ui32(v97)>>(uint(int32(2))%32)) - int32(4)
	goto L29
L40:
	;
	if v109 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v123 = base.B2i32(v68 < v121)
	if v68 < v121 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	v121 = v103 & (v103 >> (uint(int32(31)) % 32))
	goto L42
L44:
	;
	goto L45
L45:
	;
	v116 = v109 - int32(1)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v116))))
	if v118 == int32(32) {
		v109 = v116
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v121 = v109
	goto L42
L47:
	;
	goto L4
L48:
	;
	v124 = v68
	goto L50
L49:
	;
	v124 = v121
	goto L50
L50:
	;
	v125 = F_memcmp(m, v23, v76, v124)
	mBase = m.M
	if v125 != 0 {
		v128 = v125
		goto L47
	} else {
		goto L51
	}
L51:
	;
	if v68 < v121 {
		v128 = int32(-1)
		goto L47
	} else {
		goto L52
	}
L52:
	;
	v128 = base.B2i32(v121 < v68)
	goto L47
L53:
	;
	F_pfree(m, v6)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v133 != v11 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v11)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	return base.I64_extend_i32_u(base.B2i32(int32(0) <= v128))
L60:
	;
	goto L59
}
