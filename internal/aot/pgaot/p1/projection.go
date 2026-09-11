package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecAssignProjectionInfo(m *base.Module, l0 int32) {
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
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+44))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v7 = F_ExecBuildProjectionInfo(m, v3, v4, v5, l0, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v7
		return
	}
}
func F_ExecInitUpdateProjection(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v16+v17*int32(216) != l1 {
		v24 = base.I32_div_s(l1-v16, int32(216))
		v25 = v24
	} else {
		v25 = v17
	}
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v15+v25<<(uint(int32(2))%32))))
	v31 = v9 + int32(104)
	v32 = F_table_slot_create(m, v10, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v32
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v36 = F_table_slot_create(m, v35, v31)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v36
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			if v39 == int32(0) {
				F_ExecAssignExprContext(m, v9, l0)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					v46 = v45
					v47 = v44
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
					v50 = F_ExecBuildUpdateProjection(m, v48, int32(0), v29, v11, v47, v46, l0)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						v52 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)) = uint8(v52)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v50
						return
					}
				}
			} else {
				v46 = v36
				v47 = v39
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
				v50 = F_ExecBuildUpdateProjection(m, v48, int32(0), v29, v11, v47, v46, l0)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					v52 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)) = uint8(v52)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v50
					return
				}
			}
		}
	}
}
