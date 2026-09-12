package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_StandbyReleaseAllLocks(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v9 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v9 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_errmsg_internal(m, int32(154036), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[696]))
	F_hash_seq_init(m, v5+int32(12), v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	F_errfinish(m, int32(491839), int32(1111), int32(154253))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v28 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v31 = v28
	goto L13
L11:
	;
	goto L12
L12:
	;
	m.G0 = v5 + int32(32)
	return
L13:
	;
	F_StandbyReleaseXidEntryLocks(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[696]))
	v38 = F_hash_search(m, v35, v31, int32(2), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v42 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v42 != 0 {
		v31 = v42
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
}
func F_StandbyTransactionIdIsPrepared(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[102]))
	if v5 <= v2 {
		v19 = v2
		return v19
	} else {
		v9 = F_ReadTwoPhaseFile(m, l0, int32(1))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v9 == int32(0) {
				v19 = v2
				return v19
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				F_pfree(m, v9)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v19 = base.B2i32(l0 == v15)
					return v19
				}
			}
		}
	}
}
func F_standby_desc_invalidations(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	v8 = m.G0
	v10 = v8 - int32(128)
	m.G0 = v10
	if int32(0) < l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l5 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v10 + int32(128)
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+116)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = l3
	F_appendStringInfo(m, l0, int32(53632), v10+int32(112))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_appendStringInfoString(m, l0, int32(546461))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return
L8:
	;
	goto L6
L9:
	;
	v30 = int32(0)
	goto L10
L10:
	;
	v34 = l2 + v30<<(uint(int32(4))%32)
	v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v34))))
	if int32(0) <= v35 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L3
L12:
	;
	v91 = v30 + int32(1)
	if v91 != l1 {
		v30 = v91
		goto L10
	} else {
		goto L31
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v35
	F_appendStringInfo(m, l0, int32(477008), v10)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	switch v35&int32(255) - int32(250) {
	case 0:
		goto L18
	case 1:
		goto L19
	case 2:
		goto L20
	case 3:
		goto L21
	case 4:
		goto L22
	case 5:
		goto L23
	default:
		goto L17
	}
L16:
	;
	goto L12
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v35
	F_appendStringInfo(m, l0, int32(478406), v10+int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L7
	} else {
		goto L30
	}
L18:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v77
	F_appendStringInfo(m, l0, int32(55247), v10+int32(96))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L29
	}
L19:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v70
	F_appendStringInfo(m, l0, int32(40453), v10+int32(80))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L7
	} else {
		goto L28
	}
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v63
	F_appendStringInfo(m, l0, int32(55366), v10-int32(-64))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L7
	} else {
		goto L27
	}
L21:
	;
	F_appendStringInfoString(m, l0, int32(213450))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L26
	}
L22:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v53
	F_appendStringInfo(m, l0, int32(51912), v10+int32(48))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L25
	}
L23:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v46
	F_appendStringInfo(m, l0, int32(48196), v10+int32(32))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	goto L12
L25:
	;
	goto L12
L26:
	;
	goto L12
L27:
	;
	goto L12
L28:
	;
	goto L12
L29:
	;
	goto L12
L30:
	;
	goto L12
L31:
	;
	goto L11
}
