package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltxtq_rexec(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DirectFunctionCall2Coll(m, int32(5601), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_ltxtq_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v14 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v42 = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v42
	v45 = v10 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v45
	v48 = F_palloc(m, v42)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L11
	}
L6:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_errmsg(m, int32(208111), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_errdetail(m, int32(545497), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(484970), int32(614), int32(416926))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v48
	v52 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v52)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v55 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v45 + v54*v55
	F_infix_2(m, v7+v55, int32(1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_pq_begintypsend(m, v7+int32(32))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_enlargeStringInfo(m, v7+int32(32), int32(1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	v76 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v73+v74))) = uint8(v76)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v73 + v76
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v83&int32(3) == int32(0) {
		v107 = v83
		goto L17
	} else {
		goto L18
	}
L15:
	;
	F_pq_sendtext(m, v7+int32(32), v83, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L32
	}
L16:
	;
	v140 = v132 - v83
	goto L15
L17:
	;
	v111 = v107
	goto L26
L18:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v91 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v140 = int32(0)
	goto L15
L20:
	;
	goto L21
L21:
	;
	v96 = v83
	goto L22
L22:
	;
	v100 = v96 + int32(1)
	if v100&int32(3) == int32(0) {
		v107 = v100
		goto L17
	} else {
		goto L24
	}
L23:
	;
	v132 = v100
	goto L16
L24:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v105 != 0 {
		v96 = v100
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v120 = int32(-2139062144)
	if (int32(16843008)-v117|v117)&v120 == v120 {
		v111 = v111 + int32(4)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v126 = v111
	goto L29
L28:
	;
	goto L27
L29:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v130 != 0 {
		v126 = v126 + int32(1)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v132 = v126
	goto L16
L31:
	;
	goto L30
L32:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	F_pfree(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v147 = v7 + int32(32)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v150 << (uint(int32(2)) % 32)
	goto L34
L34:
	;
	m.G0 = v7 + int32(48)
	return v149
}
