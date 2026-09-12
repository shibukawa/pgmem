package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BloomNewBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int64
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(32)
	return v78
L2:
	;
	return int32(0)
L3:
	;
	if v9 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v9
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(0)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v7)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v68
	v70 = int32(8)
	v72 = int32(0)
	v75 = F_ExtendBufferedRel(m, v7+v70, v72, v72, v70)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L24
	}
L7:
	;
	v19 = F_ReadBuffer(m, l0, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v21 = F_ConditionalLockBuffer(m, v19)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v21 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v19 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	F_ReleaseBuffer(m, v19)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L21
	}
L14:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+14)))
	if v41 == int32(0) {
		v78 = v19
		goto L1
	} else {
		goto L18
	}
L15:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+(v19^int32(-1))<<(uint(int32(2))%32))))
	v40 = v32
	goto L14
L16:
	;
	goto L17
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v40 = v34 + v19<<(uint(int32(13))%32) + int32(-8192)
	goto L14
L18:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v44)+2)))
	if v46&int32(2) != 0 {
		v78 = v19
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_LockBuffer(m, v19, int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	goto L13
L21:
	;
	v55 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	if v55 != int32(-1) {
		v16 = v55
		goto L7
	} else {
		goto L23
	}
L23:
	;
	goto L8
L24:
	;
	v78 = v75
	goto L1
}
func F_bloom_get_procinfo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32)+l0)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v15 == int32(0) {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v19 = base.I32_extend16_s(l1)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+216))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+204))
		v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+6)))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v21+v23*(v19-int32(1))<<(uint(int32(2))%32)+int32(44)-int32(4))))
		if v35 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(117833860))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(250382), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(11)
						F_errdetail_internal(m, int32(654384), v8)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(497545), int32(739), int32(241756))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v40 = F_index_getprocinfo(m, v38, v19, int32(11))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v47 = v14 + int32(16)
				v48 = *(*int64)(unsafe.Add(mBase, uint32(v40)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v47))) = v48
				v50 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
				*(*int64)(unsafe.Add(mBase, uint32(v14))) = v50
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v52
				v54 = *(*int64)(unsafe.Add(mBase, uint32(v40)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v54
				*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v44
				*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(0)
				m.G0 = v8 + int32(16)
				return v14
			}
		}
	} else {
		m.G0 = v8 + int32(16)
		return v14
	}
}
