package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_statext_dependencies_serialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v10 == v2 {
		v98 = int32(16)
	} else {
		v15 = v10 & int32(3)
		v17 = l0 + int32(12)
		v18 = int32(16)
		if base.Ui32(int32(4)) <= base.Ui32(v10) {
			v24 = v18
			v26 = v2
			v30 = v2
			for {
				v34 = v17 + v26<<(uint(int32(2))%32)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
				v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35)+8)))
				v37 = int32(1)
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
				v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40)+8)))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
				v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+8)))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
				v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+8)))
				v56 = v24 + v36<<(uint(v37)%32) + v41<<(uint(v37)%32) + v46<<(uint(v37)%32) + v51<<(uint(v37)%32) + int32(40)
				v57 = int32(4)
				v58 = v26 + v57
				v60 = v30 + v57
				if v60 != v10&int32(-4) {
					v24 = v56
					v26 = v58
					v30 = v60
					continue
				} else {
					break
				}
				break
			}
			if v15 == int32(0) {
				v98 = v56
			} else {
				v65 = v56
				v67 = v58
				v74 = v65
				v76 = v67
				v81 = v2
				for {
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v17+v76<<(uint(int32(2))%32))))
					v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v85)+8)))
					v87 = int32(1)
					v91 = v74 + v86<<(uint(v87)%32) + int32(10)
					v95 = v81 + v87
					if v95 != v15 {
						v74 = v91
						v76 = v76 + v87
						v81 = v95
						continue
					} else {
						break
					}
					break
				}
				v98 = v91
			}
		} else {
			v65 = v18
			v67 = v2
			v74 = v65
			v76 = v67
			v81 = v2
			for {
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v17+v76<<(uint(int32(2))%32))))
				v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v85)+8)))
				v87 = int32(1)
				v91 = v74 + v86<<(uint(v87)%32) + int32(10)
				v95 = v81 + v87
				if v95 != v15 {
					v74 = v91
					v76 = v76 + v87
					v81 = v95
					continue
				} else {
					break
				}
				break
			}
			v98 = v91
		}
	}
	v106 = F_palloc0(m, v98)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v106))) = v98 << (uint(int32(2)) % 32)
		v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v113
		v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v106)+8)) = v115
		v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v117
		v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v119 != 0 {
			v126 = v106 + int32(16)
			v129 = int32(0)
			for {
				v137 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(12)+v129<<(uint(int32(2))%32))))
				v138 = *(*int64)(unsafe.Add(mBase, uint32(v137)))
				*(*int64)(unsafe.Add(mBase, uint32(v126))) = v138
				v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+8)))
				*(*uint16)(unsafe.Add(mBase, uint32(v126)+8)) = uint16(v140)
				v143 = v126 + int32(10)
				v144 = int32(*(*int16)(unsafe.Add(mBase, uint32(v137)+8)))
				v146 = v144 << (uint(int32(1)) % 32)
				if v146 != 0 {
					base.MemoryCopy(m, v143, v137+int32(10), v146)
				} else {
				}
				v150 = int32(*(*int16)(unsafe.Add(mBase, uint32(v137)+8)))
				v151 = int32(1)
				v155 = v129 + v151
				v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(v155) < base.Ui32(v156) {
					v126 = v143 + v150<<(uint(v151)%32)
					v129 = v155
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		return v106
	}
}
