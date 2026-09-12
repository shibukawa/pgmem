package p3

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_BloomInitMetapage(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v4 = int32(0)
	v6 = F_ReadBufferExtended(m, l0, l1, int32(-1), v4, v4)
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		F_LockBuffer(m, v6, int32(2))
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = F_GenericXLogStart(m, l0)
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v14 = F_GenericXLogRegisterBuffer(m, v11, v6, int32(1))
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					F_BloomFillMetapage(m, l0, v14)
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						F_GenericXLogFinish(m, v11)
						v19 = m.ExcPending
						if v19 != 0 {
							return
						} else {
							F_UnlockReleaseBuffer(m, v6)
							v21 = m.ExcPending
							if v21 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
