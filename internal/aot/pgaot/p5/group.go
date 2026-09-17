package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assignSortGroupRef(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	v3 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v12 == v3 {
		if l1 == int32(0) {
			v132 = int32(1)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v19 <= int32(0) {
				v132 = int32(1)
			} else {
				v22 = int32(0)
				if v22 < v19 {
					v25 = v19
				} else {
					v25 = v22
				}
				v27 = v25 & int32(3)
				v28 = int32(0)
				if int32(4) <= v19 {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v37 = v28
					v38 = int32(0)
					v39 = v3
					for {
						v48 = v33 + v39<<(uint(int32(2))%32)
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
						if base.Ui32(v37) < base.Ui32(v56) {
							v58 = v56
						} else {
							v58 = v37
						}
						if base.Ui32(v58) < base.Ui32(v54) {
							v60 = v54
						} else {
							v60 = v58
						}
						if base.Ui32(v60) < base.Ui32(v52) {
							v62 = v52
						} else {
							v62 = v60
						}
						if base.Ui32(v62) < base.Ui32(v50) {
							v64 = v50
						} else {
							v64 = v62
						}
						v65 = int32(4)
						v66 = v39 + v65
						v68 = v38 + v65
						if v68 != v25&int32(2147483644) {
							v37 = v64
							v38 = v68
							v39 = v66
							continue
						} else {
							break
						}
						break
					}
					if v27 == int32(0) {
						v110 = v64
					} else {
						v74 = v64
						v76 = v66
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v85 = int32(0)
						v87 = v74
						v89 = v76
						for {
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v83+v89<<(uint(int32(2))%32))))
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
							if base.Ui32(v87) < base.Ui32(v100) {
								v102 = v100
							} else {
								v102 = v87
							}
							v103 = int32(1)
							v106 = v85 + v103
							if v106 != v27 {
								v85 = v106
								v87 = v102
								v89 = v89 + v103
								continue
							} else {
								break
							}
							break
						}
						v110 = v102
					}
				} else {
					v74 = v28
					v76 = v3
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v85 = int32(0)
					v87 = v74
					v89 = v76
					for {
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v83+v89<<(uint(int32(2))%32))))
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
						if base.Ui32(v87) < base.Ui32(v100) {
							v102 = v100
						} else {
							v102 = v87
						}
						v103 = int32(1)
						v106 = v85 + v103
						if v106 != v27 {
							v85 = v106
							v87 = v102
							v89 = v89 + v103
							continue
						} else {
							break
						}
						break
					}
					v110 = v102
				}
				v132 = v110 + int32(1)
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v132
		v136 = v132
	} else {
		v136 = v12
	}
	return v136
}
