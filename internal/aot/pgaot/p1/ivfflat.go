package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IvfflatUpdateList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v7 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = F_ReadBufferExtended(m, l0, l5, v9, v7, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		F_LockBuffer(m, v12, int32(2))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = F_GenericXLogStart(m, l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				v20 = F_GenericXLogRegisterBuffer(m, v17, v12, int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v20+v22<<(uint(int32(2))%32))+20))
					v29 = v20 + v26&int32(32767)
					if l2 == int32(-1) {
						v40 = v7
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
						if l2 == v32 {
							v40 = v7
						} else {
							if base.B2i32(l3 != int32(-1))&base.B2i32(base.Ui32(l2) < base.Ui32(l3)) != 0 {
								v40 = v7
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = l2
								v40 = int32(1)
							}
						}
					}
					if l4 == int32(-1) {
						if v40 != 0 {
							F_GenericXLogFinish(m, v17)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v12)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							F_pfree(m, v17)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v12)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
						if l4 == v43 {
							if v40 != 0 {
								F_GenericXLogFinish(m, v17)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									F_UnlockReleaseBuffer(m, v12)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								F_pfree(m, v17)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return
								} else {
									F_UnlockReleaseBuffer(m, v12)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v29))) = l4
							F_GenericXLogFinish(m, v17)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v12)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
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
}
