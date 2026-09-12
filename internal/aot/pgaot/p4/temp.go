package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PrepareTempTablespaces(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v140 int64
	_ = v140
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[465]))
	if int32(base.Ui32(v14^int32(-1))>>(uint(int32(31))%32)) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	goto L3
L3:
	;
	if base.B2i32(v21 == int32(2)) == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[466]))
	v28 = F_pstrdup(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_pfree(m, v28)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L43
	}
L6:
	;
	return
L7:
	;
	v33 = F_SplitIdentifierString(m, v28, int32(44), v11+int32(12))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v33 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[465])) = v37
	*(*int32)(unsafe.Add(mBase, _consts[467])) = v37
	goto L14
L10:
	;
	goto L11
L11:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v58 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	goto L5
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[468])) = v37
	goto L12
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v63 = v59 << (uint(int32(2)) % 32)
	goto L18
L17:
	;
	v63 = int32(0)
	goto L18
L18:
	;
	v64 = F_MemoryContextAlloc(m, v57, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v66 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[465])) = v119
	*(*int32)(unsafe.Add(mBase, _consts[467])) = v64
	if int32(2) <= v119 {
		goto L40
	} else {
		goto L41
	}
L21:
	;
	v72 = int32(0)
	v78 = int32(0)
	goto L26
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if int32(0) < v67 {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v119 = int32(0)
	goto L20
L25:
	;
	goto L24
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v78<<(uint(int32(2))%32))))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v85 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v119 = v112
	goto L20
L28:
	;
	v116 = v78 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v116 < v117 {
		v72 = v112
		v78 = v116
		goto L26
	} else {
		goto L38
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64+v72<<(uint(int32(2))%32)))) = v104
	v112 = v72 + int32(1)
	goto L28
L30:
	;
	v104 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v90 = F_get_tablespace_oid(m, v84, int32(1))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	if v90 == int32(0) {
		v112 = v72
		goto L28
	} else {
		goto L34
	}
L34:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	if v90 == v96 {
		v104 = int32(0)
		goto L29
	} else {
		goto L35
	}
L35:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	v102 = F_object_aclcheck(m, int32(1213), v90, v100, int64(512))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	if v102 != 0 {
		v112 = v72
		goto L28
	} else {
		goto L37
	}
L37:
	;
	v104 = v90
	goto L29
L38:
	;
	goto L27
L39:
	;
	goto L5
L40:
	;
	v140 = F_pg_prng_uint64_range(m, int32(4604064), int64(0), base.I64_extend_i32_u(v119-int32(1)))
	mBase = m.M
	v142 = base.I32_wrap_i64(v140)
	goto L42
L41:
	;
	v142 = int32(0)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[468])) = v142
	goto L39
L43:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_list_free(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	goto L1
}
func F_RemoveTempRelationsCallback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[260]))
	if v8 != 0 {
		F_AbortOutOfAnyTransaction(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_StartTransactionCommand(m)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v13 = F_GetTransactionSnapshot(m)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					F_PushActiveSnapshot(m, v13)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, _consts[260]))
						*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v18
						*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(2615)
						F_performDeletion(m, v5+int32(4), int32(1), int32(29))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							F_PopActiveSnapshot(m)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_CommitTransactionCommand(m)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									m.G0 = v5 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
