package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BackendStatusShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_BackendStatusShmemInit[0]))
	if v5+int32(38) <= v2 {
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_BackendStatusShmemInit[1]))
		v12 = v11
		v13 = v2
		for {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_BackendStatusShmemInit[2]))
			*(*int32)(unsafe.Add(mBase, uint32(v16+v13*int32(408))+212)) = v12
			v24 = v13 + int32(1)
			v26 = *(*int32)(unsafe.Add(mBase, _c_F_BackendStatusShmemInit[0]))
			v28 = v26 + int32(38)
			if v24 < v28 {
				v12 = v12 - int32(-64)
				v13 = v24
				continue
			} else {
				break
			}
			break
		}
		if v28 <= int32(0) {
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_BackendStatusShmemInit[3]))
			v35 = v33
			v36 = int32(0)
			for {
				v39 = *(*int32)(unsafe.Add(mBase, _c_F_BackendStatusShmemInit[2]))
				*(*int32)(unsafe.Add(mBase, uint32(v39+v36*int32(408))+188)) = v35
				v47 = v36 + int32(1)
				v49 = *(*int32)(unsafe.Add(mBase, _c_F_BackendStatusShmemInit[0]))
				v51 = v49 + int32(38)
				if v47 < v51 {
					v35 = v35 - int32(-64)
					v36 = v47
					continue
				} else {
					break
				}
				break
			}
			if v51 <= int32(0) {
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, _c_F_BackendStatusShmemInit[4]))
				v58 = v56
				v59 = int32(0)
				for {
					v62 = *(*int32)(unsafe.Add(mBase, _c_F_BackendStatusShmemInit[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v62+v59*int32(408))+216)) = v58
					v68 = *(*int32)(unsafe.Add(mBase, _c_F_BackendStatusShmemInit[5]))
					v71 = v59 + int32(1)
					v73 = *(*int32)(unsafe.Add(mBase, _c_F_BackendStatusShmemInit[0]))
					if v71 < v73+int32(38) {
						v58 = v58 + v68
						v59 = v71
						continue
					} else {
						break
					}
					break
				}
			}
		}
	}
	return
}
