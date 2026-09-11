package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RebuildConstraintComment(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v15 = F_GetComment(m, l2, int32(2606), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if v15 != 0 {
			v18 = F_palloc0(m, int32(16))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(199)
				if l3 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(40)
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
					v26 = F_get_namespace_name(m, v25)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = F_makeString(m, v26)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v28
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
							v34 = F_pstrdup(m, v31+int32(4))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								v36 = F_makeString(m, v34)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v36
									v39 = F_pstrdup(m, l5)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										v41 = F_makeString(m, v39)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v41
											*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v41
											v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v45
											v47 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v47
											v55 = F_list_make3_impl(m, v11+int32(24), v11+int32(20), v11+int32(16))
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return
											} else {
												v79 = v55
												*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v15
												*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v79
												v83 = F_palloc0(m, int32(32))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = v18
													*(*int64)(unsafe.Add(mBase, uint32(v83))) = int64(98784247955)
													v92 = l0 + l1<<(uint(int32(2))%32) + int32(16)
													v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
													v94 = F_lappend(m, v93, v83)
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v92))) = v94
														m.G0 = v11 + int32(48)
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
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(13)
					v59 = F_copyObjectImpl(m, l4)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						v61 = F_makeTypeNameFromNameList(m, v59)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v61
							v64 = F_pstrdup(m, l5)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								v66 = F_makeString(m, v64)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v66
									*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v66
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v70
									v76 = F_list_make2_impl(m, v11+int32(12), v11+int32(8))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										v79 = v76
										*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v15
										*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v79
										v83 = F_palloc0(m, int32(32))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = v18
											*(*int64)(unsafe.Add(mBase, uint32(v83))) = int64(98784247955)
											v92 = l0 + l1<<(uint(int32(2))%32) + int32(16)
											v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
											v94 = F_lappend(m, v93, v83)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v92))) = v94
												m.G0 = v11 + int32(48)
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
		} else {
			m.G0 = v11 + int32(48)
			return
		}
	}
}
func F_get_constraint_index(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v2 = int32(0)
	v6 = F_SearchSysCache1(m, int32(19), l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v16 = v14 + v15
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+72)))
			v19 = v17 - int32(112)
			if base.Ui32(int32(8)) < base.Ui32(v19) {
				v29 = v2
			} else {
				if int32(1)<<(uint(v19)%32)&int32(289) == int32(0) {
					v29 = v2
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)+88))
					v29 = v28
				}
			}
			F_ReleaseCatCache(m, v6)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				return v29
			}
		}
	}
}
