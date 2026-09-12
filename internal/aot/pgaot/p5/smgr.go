package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_smgrDoPendingDeletes(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	goto L1
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[258]))
	if v19 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	m.G0 = v13 + int32(16)
	return
L3:
	;
	v23 = v19
	v24 = v2
	v25 = v2
	v26 = v2
	v28 = v2
	goto L4
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	if v33 < v17 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v74 <= int32(0) {
		goto L2
	} else {
		goto L28
	}
L6:
	;
	if v32 != 0 {
		v23 = v32
		v24 = v74
		v25 = v75
		v26 = v76
		v28 = v77
		goto L4
	} else {
		goto L27
	}
L7:
	;
	v74 = v24
	v75 = v25
	v76 = v26
	v77 = v23
	goto L6
L8:
	;
	goto L9
L9:
	;
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+16)))
	if l0 == v38 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v32
	goto L10
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[258])) = v32
	goto L10
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v41
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v43
	v45 = F_smgropen(m, v13, v40)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v68 = v24
	v69 = v25
	v70 = v26
	goto L16
L16:
	;
	F_pfree(m, v23)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L17
	} else {
		goto L26
	}
L17:
	;
	return
L18:
	;
	if v26 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60+v24<<(uint(int32(2))%32)))) = v45
	v68 = v24 + int32(1)
	v69 = v60
	v70 = v61
	goto L16
L20:
	;
	v51 = F_palloc(m, int32(32))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L17
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v24 < v26 {
		v60 = v25
		v61 = v26
		goto L19
	} else {
		goto L24
	}
L23:
	;
	v60 = v51
	v61 = int32(8)
	goto L19
L24:
	;
	v56 = F_repalloc(m, v25, v26<<(uint(int32(3))%32))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v60 = v56
	v61 = v26 << (uint(int32(1)) % 32)
	goto L19
L26:
	;
	v74 = v68
	v75 = v69
	v76 = v70
	v77 = v28
	goto L6
L27:
	;
	goto L5
L28:
	;
	v81 = int32(0)
	F_smgrdounlinkall(m, v75, v74, v81)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	v86 = v81
	goto L30
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v75+v86<<(uint(int32(2))%32))))
	F_smgrclose(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L17
	} else {
		goto L32
	}
L31:
	;
	F_pfree(m, v75)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L17
	} else {
		goto L34
	}
L32:
	;
	v102 = v86 + int32(1)
	if v102 != v74 {
		v86 = v102
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L2
}
func F_smgr_aio_reopen(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = l0 + int32(104)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)))
	if v15&int32(1) != 0 {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v19 = v18
	} else {
		v19 = int32(-1)
	}
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v20
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v22
	v24 = F_smgropen(m, v8, v19)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return
	} else {
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
		v27 = int32(1)
		if base.Ui32(v26-v27) <= base.Ui32(v27) {
			v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+20)))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v35*int32(80))+uint32(_consts[794])))
			v41 = m.T0[v40].(func(*base.Module, int32, int32, int32, int32) int32)(m, v24, v31, v32, v8+int32(12))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0+int32(88)))) = v41
				m.G0 = v8 + int32(16)
				return
			}
		} else {
			m.G0 = v8 + int32(16)
			return
		}
	}
}
