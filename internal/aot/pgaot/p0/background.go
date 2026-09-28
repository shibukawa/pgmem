package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BackgroundWorkerShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	v2 = int32(0)
	v5 = int32(_a_F_BackgroundWorkerShmemInit_0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerShmemInit[0]))
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerShmemInit[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerShmemInit[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerShmemInit[2]))
	if base.B2i32(v15 == v2)|base.B2i32(v15 == int32(_a_F_BackgroundWorkerShmemInit_1)) == v2 {
		v23 = v2
		v25 = v15
		for {
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerShmemInit[0]))
			v30 = v27 + v23*int32(1488)
			*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = int32(-1)
			v35 = int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)) = uint16(v35)
			v37 = int32(32)
			*(*int32)(unsafe.Add(mBase, uint32(v25-v37))) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v25-int32(8)))) = v23
			base.MemoryCopy(m, v30+v37, v25-int32(1496), int32(1472))
			v51 = v23 + v35
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
			if v52 != int32(_a_F_BackgroundWorkerShmemInit_1) {
				v23 = v51
				v25 = v52
				continue
			} else {
				break
			}
			break
		}
		v56 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerShmemInit[1]))
		v57 = v51
		v58 = v56
	} else {
		v57 = v2
		v58 = v8
	}
	if v57 < v58 {
		v61 = v57
		for {
			v65 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerShmemInit[0]))
			v69 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v65+v61*int32(1488))+16)) = uint8(v69)
			v72 = v61 + int32(1)
			v74 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerShmemInit[1]))
			if v72 < v74 {
				v61 = v72
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
