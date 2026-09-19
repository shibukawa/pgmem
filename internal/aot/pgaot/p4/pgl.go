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
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	v1 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*uint8)(unsafe.Add(mBase, _c_F_pgl_sendConnData[0])) = uint8(v1)
	F_pq_beginmessage(m, v6, int32(82))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		F_enlargeStringInfo(m, v6, int32(4))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*int32)(unsafe.Add(mBase, uint32(v17+v18))) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v17 + int32(4)
			F_pq_endmessage(m, v6)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_BeginReportingGUCOptions(m)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_pgstat_report_connect(m)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_pq_beginmessage(m, v6, int32(75))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_sendConnData[1]))
							F_enlargeStringInfo(m, v6, int32(4))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
								v46 = int32(16711935)
								*(*int32)(unsafe.Add(mBase, uint32(v41+v42))) = base.I32_rotr(v37, int32(24))&v46 | base.I32_rotr(v37&v46, int32(8))
								v54 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v41 + v54
								F_enlargeStringInfo(m, v6, v54)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
									*(*int32)(unsafe.Add(mBase, uint32(v60+v61))) = int32(-528071424)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v60 + int32(4)
									F_pq_endmessage(m, v6)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										F_ReadyForQuery(m, int32(2))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return
										} else {
											m.G0 = v6 + int32(16)
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
