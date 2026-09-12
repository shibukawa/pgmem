package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_QueuePartitionConstraintValidation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	v4 = l3
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = F_PartConstraintImpliedByRelConstraint(m, l1, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return
L2:
	;
	return
L3:
	;
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+119)))
	switch v54 - int32(112) {
	case 0:
		goto L17
	default:
		goto L1
	case 2:
		goto L18
	}
L7:
	;
	if v4 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v19 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	if v19 == int32(0) {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v25 + int32(4)
	F_errmsg_internal(m, int32(128208), v13+int32(16))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(520068), int32(20190), int32(278626))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	goto L1
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41 + int32(4)
	F_errmsg_internal(m, int32(128117), v13)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(520068), int32(20194), int32(278626))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	goto L1
L17:
	;
	v133 = F_RelationGetPartitionDesc(m, l1, int32(1))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L30
	}
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v58 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+104)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v123)+100)) = v129
	goto L1
L20:
	;
	v96 = F_palloc0(m, int32(140))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L2
	} else {
		goto L27
	}
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v61 <= int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v70 = int32(0)
	goto L23
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v64+v70<<(uint(int32(2))%32))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v80 == v57 {
		v123 = v79
		goto L19
	} else {
		goto L25
	}
L24:
	;
	goto L20
L25:
	;
	v83 = v70 + int32(1)
	if v61 != v83 {
		v70 = v83
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v57
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+4)) = uint8(v102)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v105 = F_CreateTupleDescCopyConstr(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v105
	*(*int64)(unsafe.Add(mBase, uint32(v96)+88)) = int64(0)
	v110 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+84)) = uint8(v110)
	v112 = int32(28672)
	*(*uint16)(unsafe.Add(mBase, uint32(v96)+96)) = uint16(v112)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v115 = F_lappend(m, v114, v96)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v115
	v123 = v96
	goto L19
L30:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	if v135 <= int32(0) {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v143 = int32(0)
	goto L32
L32:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148+v143<<(uint(int32(2))%32))))
	v154 = F_table_open(m, v152, int32(8))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L34
	}
L33:
	;
	goto L1
L34:
	;
	v157 = F_map_partition_varattnos(m, l2, int32(1), v154, l1)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	F_QueuePartitionConstraintValidation(m, l0, v154, v157, v4)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	F_sequence_close(m, v154, int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v165 = v143 + int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	if v165 < v166 {
		v143 = v165
		goto L32
	} else {
		goto L38
	}
L38:
	;
	goto L33
}
func F_quote_ident(m *base.Module, l0 int32) int32 {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_text_to_cstring(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			v9 = F_quote_identifier(m, v7)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int32(0)
			} else {
				v11 = F_cstring_to_text(m, v9)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return int32(0)
				} else {
					return v11
				}
			}
		}
	}
}
