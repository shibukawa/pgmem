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
	var v5 int32
	_ = v5
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	v4 = l3
	v5 = int32(0)
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
	F_errmsg_internal(m, int32(_a_F_QueuePartitionConstraintValidation_0), v13+int32(16))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_QueuePartitionConstraintValidation_1), int32(_a_F_QueuePartitionConstraintValidation_2), int32(_a_F_QueuePartitionConstraintValidation_3))
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
	F_errmsg_internal(m, int32(_a_F_QueuePartitionConstraintValidation_4), v13)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_QueuePartitionConstraintValidation_1), int32(_a_F_QueuePartitionConstraintValidation_5), int32(_a_F_QueuePartitionConstraintValidation_3))
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
	v132 = F_RelationGetPartitionDesc(m, l1, int32(1))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
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
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+104)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+100)) = v128
	goto L1
L20:
	;
	v95 = F_palloc0(m, int32(140))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
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
	v70 = v5
	goto L23
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v64+v70<<(uint(int32(2))%32))))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v79 == v57 {
		v121 = v78
		goto L19
	} else {
		goto L25
	}
L24:
	;
	goto L20
L25:
	;
	v82 = v70 + int32(1)
	if v61 != v82 {
		v70 = v82
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v57
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+4)) = uint8(v101)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v104 = F_CreateTupleDescCopyConstr(m, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v95)+88)) = int64(0)
	v109 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+84)) = uint8(v109)
	v111 = int32(_a_F_QueuePartitionConstraintValidation_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v95)+96)) = uint16(v111)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v114 = F_lappend(m, v113, v95)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v114
	v121 = v95
	goto L19
L30:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v134 <= int32(0) {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v141 = v5
	goto L32
L32:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147+v141<<(uint(int32(2))%32))))
	v153 = F_table_open(m, v151, int32(8))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L2
	} else {
		goto L34
	}
L33:
	;
	goto L1
L34:
	;
	v156 = F_map_partition_varattnos(m, l2, int32(1), v153, l1)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	F_QueuePartitionConstraintValidation(m, l0, v153, v156, v4)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	F_relation_close(m, v153, int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v164 = v141 + int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v164 < v165 {
		v141 = v164
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
