package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_attribute_aclmask_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v75 int64
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int64
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int64
	_ = v136
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v16 = base.I64_extend_i32_u(l0)
	v18 = F_SearchSysCache2(m, int32(7), v16, base.I64_extend_i32_s(l1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		if v18 == int32(0) {
			if l4 != 0 {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v24)
				v136 = int64(0)
				m.G0 = v13 + int32(48)
				return v136
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50360452))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
						F_errmsg(m, int32(_a_F_pg_attribute_aclmask_ext_0), v13)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_attribute_aclmask_ext_1), int32(3181), int32(_a_F_pg_attribute_aclmask_ext_2))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+22)))
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v45)+91)))
			if v47 == int32(1) {
				if l4 != 0 {
					v50 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v50)
					F_ReleaseCatCache(m, v18)
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int64(0)
					} else {
						v136 = int64(0)
						m.G0 = v13 + int32(48)
						return v136
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50360452))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
							F_errmsg(m, int32(_a_F_pg_attribute_aclmask_ext_0), v13+int32(16))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_pg_attribute_aclmask_ext_1), int32(3200), int32(_a_F_pg_attribute_aclmask_ext_2))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				v75 = F_SysCacheGetAttr(m, int32(7), v18, int32(22), v13+int32(47))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int64(0)
				} else {
					v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)))
					if v77 != int32(1) {
						v85 = F_SearchSysCache1(m, int32(57), v16)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int64(0)
						} else {
							if v85 == int32(0) {
								F_ReleaseCatCache(m, v18)
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int64(0)
								} else {
									if l4 != 0 {
										v91 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v91)
										v136 = int64(0)
										m.G0 = v13 + int32(48)
										return v136
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(16908420))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l0
												F_errmsg(m, int32(_a_F_pg_attribute_aclmask_ext_3), v13+int32(32))
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_pg_attribute_aclmask_ext_1), int32(3238), int32(_a_F_pg_attribute_aclmask_ext_2))
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									}
								}
							} else {
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+22)))
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v112+v113)+80))
								F_ReleaseCatCache(m, v85)
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return int64(0)
								} else {
									v118 = base.I32_wrap_i64(v75)
									v119 = F_pg_detoast_datum(m, v118)
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int64(0)
									} else {
										v122 = F_aclmask(m, v119, l2, v115, l3, int32(1))
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return int64(0)
										} else {
											v124 = int32(0)
											if base.B2i32(v119 == v124)|base.B2i32(v119 == v118) == v124 {
												F_pfree(m, v119)
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return int64(0)
												} else {
													F_ReleaseCatCache(m, v18)
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int64(0)
													} else {
														v136 = v122
														m.G0 = v13 + int32(48)
														return v136
													}
												}
											} else {
												F_ReleaseCatCache(m, v18)
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return int64(0)
												} else {
													v136 = v122
													m.G0 = v13 + int32(48)
													return v136
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_ReleaseCatCache(m, v18)
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int64(0)
						} else {
							v136 = int64(0)
							m.G0 = v13 + int32(48)
							return v136
						}
					}
				}
			}
		}
	}
}
