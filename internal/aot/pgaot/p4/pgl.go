package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgl_pq_flush(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_pq_flush[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	v4 = m.T0[v3].(func(*base.Module) int32)(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_pgl_sendConnData(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, _c_F_pgl_sendConnData[0])) = uint8(v1)
	F_pq_beginmessage(m, v7, int32(82))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		F_enlargeStringInfo(m, v7, int32(4))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			*(*int32)(unsafe.Add(mBase, uint32(v18+v19))) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18 + int32(4)
			F_pq_endmessage(m, v7)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				F_BeginReportingGUCOptions(m)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_pgstat_report_connect(m)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						F_pq_beginmessage(m, v7, int32(75))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_sendConnData[1]))
							F_enlargeStringInfo(m, v7, int32(4))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
								v45 = int32(24)
								v47 = int32(_a_F_pgl_sendConnData_0)
								v49 = int32(8)
								*(*int32)(unsafe.Add(mBase, uint32(v42+v43))) = v38<<(uint(v45)%32) | v38&v47<<(uint(v49)%32) | (int32(base.Ui32(v38)>>(uint(v49)%32))&v47 | int32(base.Ui32(v38)>>(uint(v45)%32)))
								v61 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v42 + v61
								F_enlargeStringInfo(m, v7, v61)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
									*(*int32)(unsafe.Add(mBase, uint32(v67+v68))) = int32(1087521792)
									*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v67 + int32(4)
									F_pq_endmessage(m, v7)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										F_ReadyForQuery(m, int32(2))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pgl_shmdt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_shmdt[0]))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgl_shmdt[1])) = int32(28)
	return int32(-1)
L2:
	;
	v8 = v4
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if l0 != v9 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v11 != 0 {
		v8 = v11
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	goto L4
L8:
	;
	goto L1
}
