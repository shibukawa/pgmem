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
	var v36 int32
	_ = v36
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
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
					v29 = v20 + v26&int32(_a_F_IvfflatUpdateList_0)
					if l2 == int32(-1) {
						v41 = v7
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
						if base.B2i32(l3 != int32(-1))&base.B2i32(base.Ui32(l2) < base.Ui32(l3))|base.B2i32(l2 == v36) != 0 {
							v41 = v7
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = l2
							v41 = int32(1)
						}
					}
					if l4 == int32(-1) {
						if v41 != 0 {
							F_GenericXLogFinish(m, v17)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v12)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							F_pfree(m, v17)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v12)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
						if l4 == v44 {
							if v41 != 0 {
								F_GenericXLogFinish(m, v17)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									F_UnlockReleaseBuffer(m, v12)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								F_pfree(m, v17)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									F_UnlockReleaseBuffer(m, v12)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
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
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v12)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
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
