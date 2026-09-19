package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ProcessClientReadInterrupt(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientReadInterrupt[0]))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessClientReadInterrupt[1])))
	if v6 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessClientReadInterrupt[0])) = v4
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientReadInterrupt[2]))
	if v8 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientReadInterrupt[3]))
	if v25 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientReadInterrupt[4]))
	if v12 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	F_ProcessCatchupInterrupt(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientReadInterrupt[5]))
	if v16 == int32(0) {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	F_ProcessNotifyInterrupt(m, int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessClientReadInterrupt[0])) = v4
	return
L16:
	;
	if l0 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientReadInterrupt[2]))
	if v29 == int32(0) {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientReadInterrupt[6]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v38 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessClientReadInterrupt[0])) = v4
	return
L22:
	;
	goto L1
L23:
	;
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v41 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v44 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientReadInterrupt[7]))
	if v48 == v44 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v50 = m.G0
	v52 = v50 - int32(16)
	m.G0 = v52
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientReadInterrupt[8]))
	if v55 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v78 = F_pgmem_kill(m, v44, int32(23))
	mBase = m.M
	goto L23
L30:
	;
	m.G0 = v52 + int32(16)
	goto L22
L31:
	;
	v58 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+15)) = uint8(v58)
	goto L32
L32:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientReadInterrupt[9]))
	v66 = F_write(m, v62, v52+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v66 {
		goto L30
	} else {
		goto L34
	}
L33:
	;
	goto L30
L34:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientReadInterrupt[0]))
	if v70 == int32(27) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
}
func F_assign_client_encoding(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_assign_client_encoding[0]))
	if int32(0) <= v8 {
		m.G0 = v5 + int32(16)
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v12 = F_SetClientEncoding(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if int32(0) <= v12 {
				m.G0 = v5 + int32(16)
				return
			} else {
				v18 = F_errstart(m, int32(15), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					if v18 == int32(0) {
						m.G0 = v5 + int32(16)
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = v11
						F_errmsg_internal(m, int32(_a_F_assign_client_encoding_0), v5)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_assign_client_encoding_1), int32(799), int32(_a_F_assign_client_encoding_2))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
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
}
