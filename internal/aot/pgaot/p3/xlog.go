package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SendXlogRecPtrResult(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	v1 = l0
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+22)) = uint16(v3)
	v13 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v16 = F_CreateTemplateTupleDesc(m, int32(2))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_TupleDescInitBuiltinEntry(m, v16, int32(1), int32(215595), int32(25))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_TupleDescInitBuiltinEntry(m, v16, int32(2), int32(333078), int32(20))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v29 = F_begin_tup_output_tupdesc(m, v13, v16, int32(1646164))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)) = uint32(v1)
						v33 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
						*(*uint32)(unsafe.Add(mBase, uint32(v8))) = uint32(v33)
						v36 = F_psprintf(m, int32(537209), v8)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v38 = F_cstring_to_text(m, v36)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v38
								v42 = F_Int64GetDatum(m, base.I64_extend_i32_u(l1))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v42
									F_do_tup_output(m, v29, v8+int32(24), v8+int32(22))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										F_end_tup_output(m, v29)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return
										} else {
											F_pq_puttextmessage(m)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return
											} else {
												m.G0 = v8 + int32(32)
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
}
