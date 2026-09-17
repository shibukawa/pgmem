package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SnapBuildProcessChange(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 <= v4 {
		v92 = v4
		return v92
	} else {
		if v8 == int32(1) {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v13))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l1)) == int32(0) {
				v25 = base.B2i32(base.Ui32(l1) < base.Ui32(v13))
			} else {
				v25 = int32(base.Ui32(l1-v13) >> (uint(int32(31)) % 32))
			}
			if v25 != 0 {
				v92 = v4
				return v92
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v28 = F_ReorderBufferXidHasBaseSnapshot(m, v27, l1)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					if v28 != 0 {
						v92 = int32(1)
						return v92
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v32 == int32(0) {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							v41 = F_MemoryContextAllocZero(m, v35, v36<<(uint(int32(2))%32)+int32(76))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(5)
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v45
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v49 = v41 + int32(72)
								*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v49
								*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v47
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v52
								v55 = v52 << (uint(int32(2)) % 32)
								if v55 != 0 {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
									base.MemoryCopy(m, v49, v56, v55)
								} else {
								}
								F_pg_qsort(m, v49, v52, int32(4), int32(185))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v62 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v41)+64)) = v62
									*(*int64)(unsafe.Add(mBase, uint32(v41)+44)) = v62
									v66 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v66
									*(*int64)(unsafe.Add(mBase, uint32(v41)+20)) = v62
									*(*int32)(unsafe.Add(mBase, uint32(v41)+27)) = v66
									*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v41
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v41)+44))
									*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v73 + int32(1)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v78 = v77
									v82 = int32(1)
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
									*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v83 + v82
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									F_ReorderBufferSetBaseSnapshot(m, v87, l1, l2, v88)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int32(0)
									} else {
										v92 = v82
										return v92
									}
								}
							}
						} else {
							v78 = v32
							v82 = int32(1)
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
							*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v83 + v82
							v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							F_ReorderBufferSetBaseSnapshot(m, v87, l1, l2, v88)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								v92 = v82
								return v92
							}
						}
					}
				}
			}
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v28 = F_ReorderBufferXidHasBaseSnapshot(m, v27, l1)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				if v28 != 0 {
					v92 = int32(1)
					return v92
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v32 == int32(0) {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
						v41 = F_MemoryContextAllocZero(m, v35, v36<<(uint(int32(2))%32)+int32(76))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(5)
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v45
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v49 = v41 + int32(72)
							*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v49
							*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v47
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v52
							v55 = v52 << (uint(int32(2)) % 32)
							if v55 != 0 {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
								base.MemoryCopy(m, v49, v56, v55)
							} else {
							}
							F_pg_qsort(m, v49, v52, int32(4), int32(185))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v62 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v41)+64)) = v62
								*(*int64)(unsafe.Add(mBase, uint32(v41)+44)) = v62
								v66 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v41)+20)) = v62
								*(*int32)(unsafe.Add(mBase, uint32(v41)+27)) = v66
								*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v41
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v41)+44))
								*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v73 + int32(1)
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								v78 = v77
								v82 = int32(1)
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
								*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v83 + v82
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								F_ReorderBufferSetBaseSnapshot(m, v87, l1, l2, v88)
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int32(0)
								} else {
									v92 = v82
									return v92
								}
							}
						}
					} else {
						v78 = v32
						v82 = int32(1)
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
						*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v83 + v82
						v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						F_ReorderBufferSetBaseSnapshot(m, v87, l1, l2, v88)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							v92 = v82
							return v92
						}
					}
				}
			}
		}
	}
}
