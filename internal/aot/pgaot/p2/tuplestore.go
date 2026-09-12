package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tuplestore_advance(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_tuplestore_gettuple(m, l0, l1, v6+int32(15))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			m.G0 = v6 + int32(16)
			return base.B2i32(v10 != int32(0))
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
			if v16 != int32(1) {
				m.G0 = v6 + int32(16)
				return base.B2i32(v10 != int32(0))
			} else {
				F_pfree(m, v10)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return base.B2i32(v10 != int32(0))
				}
			}
		}
	}
}
func F_tuplestore_alloc_read_pointer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v6 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
		if v9 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v16 < v17 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				v32 = v19
				v33 = v16
				v34 = int32(24)
				v36 = v32 + v33*v34
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
				*(*int64)(unsafe.Add(mBase, uint32(v36))) = v37
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v41
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				*(*int32)(unsafe.Add(mBase, uint32(v43+v44*v34))) = l1
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49 | l1
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v52 + int32(1)
				return v52
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				v23 = F_repalloc(m, v20, v17*int32(48))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v17 << (uint(int32(1)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v23
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					v32 = v23
					v33 = v31
					v34 = int32(24)
					v36 = v32 + v33*v34
					v37 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
					*(*int64)(unsafe.Add(mBase, uint32(v36))) = v37
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39
					v41 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v41
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					*(*int32)(unsafe.Add(mBase, uint32(v43+v44*v34))) = l1
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49 | l1
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v52 + int32(1)
					return v52
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v12|l1 != v12 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(154171), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(491070), int32(401), int32(212655))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				if v16 < v17 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
					v32 = v19
					v33 = v16
					v34 = int32(24)
					v36 = v32 + v33*v34
					v37 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
					*(*int64)(unsafe.Add(mBase, uint32(v36))) = v37
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39
					v41 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v41
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					*(*int32)(unsafe.Add(mBase, uint32(v43+v44*v34))) = l1
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49 | l1
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v52 + int32(1)
					return v52
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
					v23 = F_repalloc(m, v20, v17*int32(48))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v17 << (uint(int32(1)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v23
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						v32 = v23
						v33 = v31
						v34 = int32(24)
						v36 = v32 + v33*v34
						v37 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
						*(*int64)(unsafe.Add(mBase, uint32(v36))) = v37
						v39 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39
						v41 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v41
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						*(*int32)(unsafe.Add(mBase, uint32(v43+v44*v34))) = l1
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49 | l1
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v52 + int32(1)
						return v52
					}
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v12|l1 != v12 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(154171), int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491070), int32(401), int32(212655))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v16 < v17 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				v32 = v19
				v33 = v16
				v34 = int32(24)
				v36 = v32 + v33*v34
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
				*(*int64)(unsafe.Add(mBase, uint32(v36))) = v37
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v41
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				*(*int32)(unsafe.Add(mBase, uint32(v43+v44*v34))) = l1
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49 | l1
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v52 + int32(1)
				return v52
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				v23 = F_repalloc(m, v20, v17*int32(48))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v17 << (uint(int32(1)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v23
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					v32 = v23
					v33 = v31
					v34 = int32(24)
					v36 = v32 + v33*v34
					v37 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
					*(*int64)(unsafe.Add(mBase, uint32(v36))) = v37
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39
					v41 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v41
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					*(*int32)(unsafe.Add(mBase, uint32(v43+v44*v34))) = l1
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49 | l1
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v52 + int32(1)
					return v52
				}
			}
		}
	}
}
func F_tuplestore_copy_read_pointer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	if l1 == l2 {
		return
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
		v17 = int32(24)
		v19 = v16 + l1*v17
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
		v23 = v16 + l2*v17
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		v25 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
		*(*int64)(unsafe.Add(mBase, uint32(v23))) = v25
		v27 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v27
		v29 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v29
		if v20 != v24 {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			if v34 < int32(2) {
				v134 = v33
			} else {
				v38 = v34 - int32(1)
				v39 = int32(3)
				v40 = v38 & v39
				if base.Ui32(v34-int32(2)) < base.Ui32(v39) {
					v91 = int32(1)
					v93 = v33
				} else {
					v58 = int32(1)
					v60 = v33
					v64 = int32(0)
					for {
						v70 = v58 * int32(24)
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(72)+v70)))
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+(v32+int32(48)))))
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v70+(v32+int32(24)))))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v70+v32)))
						v82 = v72 | (v74 | (v76 | (v78 | v60)))
						v83 = int32(4)
						v84 = v58 + v83
						v86 = v64 + v83
						if v86 != v38&int32(-4) {
							v58 = v84
							v60 = v82
							v64 = v86
							continue
						} else {
							break
						}
						break
					}
					v91 = v84
					v93 = v82
				}
				if v40 == int32(0) {
					v134 = v93
				} else {
					v108 = v91
					v109 = int32(0)
					v110 = v93
					for {
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v32+v108*int32(24))))
						v123 = v122 | v110
						v124 = int32(1)
						v127 = v109 + v124
						if v127 != v40 {
							v108 = v108 + v124
							v109 = v127
							v110 = v123
							continue
						} else {
							break
						}
						break
					}
					v134 = v123
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v134
		} else {
		}
		v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(v158) < base.Ui32(int32(2)) {
			return
		} else {
			if v158 != int32(2) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v243 = m.ExcPending
				if v243 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(347462), int32(0))
					mBase = m.M
					v247 = m.ExcPending
					if v247 != 0 {
						return
					} else {
						F_errfinish(m, int32(491070), int32(1394), int32(212595))
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				if v163 == l2 {
					v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
					if v166 == int32(1) {
						v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
						v170 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
						v172 = F_BufFileSeek(m, v165, v169, v170, int32(0))
						mBase = m.M
						v173 = m.ExcPending
						if v173 != 0 {
							return
						} else {
							if v172 == int32(0) {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v179 = m.ExcPending
								if v179 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v181 = m.ExcPending
									if v181 != 0 {
										return
									} else {
										F_errmsg(m, int32(381221), int32(0))
										mBase = m.M
										v185 = m.ExcPending
										if v185 != 0 {
											return
										} else {
											F_errfinish(m, int32(491070), int32(1373), int32(212595))
											mBase = m.M
											v190 = m.ExcPending
											if v190 != 0 {
												return
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
						v191 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
						v192 = *(*int64)(unsafe.Add(mBase, uint32(v23)+16))
						v194 = F_BufFileSeek(m, v165, v191, v192, int32(0))
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
							return
						} else {
							if v194 == int32(0) {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v201 = m.ExcPending
								if v201 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return
									} else {
										F_errmsg(m, int32(381221), int32(0))
										mBase = m.M
										v207 = m.ExcPending
										if v207 != 0 {
											return
										} else {
											F_errfinish(m, int32(491070), int32(1382), int32(212595))
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return
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
					}
				} else {
					if l1 != v163 {
					} else {
						v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
						if v214 != 0 {
						} else {
							v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v220 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v23+int32(12)))) = v220
							v222 = *(*int64)(unsafe.Add(mBase, uint32(v215)+32))
							v223 = int64(*(*int32)(unsafe.Add(mBase, uint32(v215)+40)))
							*(*int64)(unsafe.Add(mBase, uint32(v23+int32(16)))) = v222 + v223
						}
					}
					return
				}
			}
		}
	}
}
func F_tuplestore_get_stats(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 == int32(0) {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v12 = v10 - v11
		if v12 < v6 {
			v14 = v6
		} else {
			v14 = v12
		}
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v14
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)))
		if v18&int32(1) != 0 {
			v21 = int32(310126)
		} else {
			v21 = int32(14035)
		}
		v37 = v21
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v37
		v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int64)(unsafe.Add(mBase, uint32(l2))) = v39
		return
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v23 = F_BufFileSize(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			if v23 < v6 {
				v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
				v30 = v26
				v31 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v31)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v30
				v37 = int32(310126)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v37
				v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int64)(unsafe.Add(mBase, uint32(l2))) = v39
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v28 = F_BufFileSize(m, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v30 = v28
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v31)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v30
					v37 = int32(310126)
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v37
					v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = v39
					return
				}
			}
		}
	}
}
func F_tuplestore_gettuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v15 = v11 + v12*int32(24)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v16 {
	case 0:
		v17 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v17)
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
		if l1 != 0 {
			if v19&int32(1) != 0 {
				v206 = v4
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
				if v23 <= v22 {
					v201 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v201)
					v206 = v4
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v22 + int32(1)
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v25+v22<<(uint(int32(2))%32))))
					v206 = v32
				}
			}
		} else {
			if v19&int32(1) != 0 {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
				v36 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v36)
				v44 = v35
				*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v44
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				if v44 <= v46 {
					v206 = int32(0)
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+v44<<(uint(int32(2))%32)-int32(4))))
					v206 = v55
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				if v38 <= v39 {
					v206 = int32(0)
				} else {
					v44 = v38 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v44
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					if v44 <= v46 {
						v206 = int32(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+v44<<(uint(int32(2))%32)-int32(4))))
						v206 = v55
					}
				}
			}
		}
		m.G0 = v9 + int32(16)
		return v206
	case 1:
		if l1 == int32(0) {
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
			*(*int32)(unsafe.Add(mBase, uint32(l0+int32(108)))) = v68
			v70 = *(*int64)(unsafe.Add(mBase, uint32(v63)+32))
			v71 = int64(*(*int32)(unsafe.Add(mBase, uint32(v63)+40)))
			*(*int64)(unsafe.Add(mBase, uint32(l0+int32(112)))) = v70 + v71
			v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
			if v74 == int32(0) {
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
				v79 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
				v81 = F_BufFileSeek(m, v77, v78, v79, int32(0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					if v81 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v215 = m.ExcPending
						if v215 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v217 = m.ExcPending
							if v217 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(381221), int32(0))
								mBase = m.M
								v221 = m.ExcPending
								if v221 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(491070), int32(1025), int32(377494))
									mBase = m.M
									v226 = m.ExcPending
									if v226 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2)
						v87 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v87)
						if l1 != 0 {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v90 = int32(4)
							v94 = F_BufFileReadMaybeEOF(m, v89, v9+v90, v90, int32(1))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								if v94 == int32(0) {
									v105 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
									v206 = int32(0)
									m.G0 = v9 + int32(16)
									return v206
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
									if v98 == int32(0) {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
										v206 = int32(0)
										m.G0 = v9 + int32(16)
										return v206
									} else {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
										v102 = m.T0[v101].(func(*base.Module, int32, int32) int32)(m, l0, v98)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int32(0)
										} else {
											v206 = v102
											m.G0 = v9 + int32(16)
											return v206
										}
									}
								}
							}
						} else {
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v112 = F_BufFileSeek(m, v108, int32(0), int64(-4), int32(1))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int32(0)
							} else {
								if v112 != 0 {
									v114 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v114)
									v206 = v4
									m.G0 = v9 + int32(16)
									return v206
								} else {
									v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v121 = F_BufFileReadMaybeEOF(m, v116, v9+int32(8), int32(4), int32(0))
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int32(0)
									} else {
										v123 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
										if v121 != 0 {
											v125 = v123
										} else {
											v125 = int32(0)
										}
										v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
										if v126 == int32(1) {
											v129 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v129)
											v175 = v125
											v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v178 = int32(0)
											v183 = F_BufFileSeek(m, v177, v178, base.I64_extend_i32_s(v178-v175), int32(1))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return int32(0)
											} else {
												if v183 != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v230 = m.ExcPending
													if v230 != 0 {
														return int32(0)
													} else {
														F_errcode_for_file_access(m)
														mBase = m.M
														v232 = m.ExcPending
														if v232 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(381221), int32(0))
															mBase = m.M
															v236 = m.ExcPending
															if v236 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(491070), int32(1106), int32(377494))
																mBase = m.M
																v241 = m.ExcPending
																if v241 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
													v186 = m.T0[v185].(func(*base.Module, int32, int32) int32)(m, l0, v175)
													mBase = m.M
													v187 = m.ExcPending
													if v187 != 0 {
														return int32(0)
													} else {
														v206 = v186
														m.G0 = v9 + int32(16)
														return v206
													}
												}
											}
										} else {
											v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v137 = F_BufFileSeek(m, v131, int32(0), base.I64_extend_i32_s(int32(-8)-v125), int32(1))
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return int32(0)
											} else {
												if v137 != 0 {
													v139 = int32(0)
													v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													v146 = F_BufFileSeek(m, v140, v139, base.I64_extend_i32_s(int32(-4)-v125), int32(1))
													mBase = m.M
													v147 = m.ExcPending
													if v147 != 0 {
														return int32(0)
													} else {
														if v146 == int32(0) {
															v206 = v139
															m.G0 = v9 + int32(16)
															return v206
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																F_errcode_for_file_access(m)
																mBase = m.M
																v155 = m.ExcPending
																if v155 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(381221), int32(0))
																	mBase = m.M
																	v159 = m.ExcPending
																	if v159 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(491070), int32(1089), int32(377494))
																		mBase = m.M
																		v164 = m.ExcPending
																		if v164 != 0 {
																			return int32(0)
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
													v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													v170 = F_BufFileReadMaybeEOF(m, v165, v9+int32(12), int32(4), int32(0))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int32(0)
													} else {
														v172 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
														if v170 != 0 {
															v174 = v172
														} else {
															v174 = int32(0)
														}
														v175 = v174
														v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
														v178 = int32(0)
														v183 = F_BufFileSeek(m, v177, v178, base.I64_extend_i32_s(v178-v175), int32(1))
														mBase = m.M
														v184 = m.ExcPending
														if v184 != 0 {
															return int32(0)
														} else {
															if v183 != 0 {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v230 = m.ExcPending
																if v230 != 0 {
																	return int32(0)
																} else {
																	F_errcode_for_file_access(m)
																	mBase = m.M
																	v232 = m.ExcPending
																	if v232 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(381221), int32(0))
																		mBase = m.M
																		v236 = m.ExcPending
																		if v236 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(491070), int32(1106), int32(377494))
																			mBase = m.M
																			v241 = m.ExcPending
																			if v241 != 0 {
																				return int32(0)
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		}
																	}
																}
															} else {
																v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
																v186 = m.T0[v185].(func(*base.Module, int32, int32) int32)(m, l0, v175)
																mBase = m.M
																v187 = m.ExcPending
																if v187 != 0 {
																	return int32(0)
																} else {
																	v206 = v186
																	m.G0 = v9 + int32(16)
																	return v206
																}
															}
														}
													}
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
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2)
				v87 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v87)
				if l1 != 0 {
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v90 = int32(4)
					v94 = F_BufFileReadMaybeEOF(m, v89, v9+v90, v90, int32(1))
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						if v94 == int32(0) {
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
							v206 = int32(0)
							m.G0 = v9 + int32(16)
							return v206
						} else {
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							if v98 == int32(0) {
								v105 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
								v206 = int32(0)
								m.G0 = v9 + int32(16)
								return v206
							} else {
								v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
								v102 = m.T0[v101].(func(*base.Module, int32, int32) int32)(m, l0, v98)
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int32(0)
								} else {
									v206 = v102
									m.G0 = v9 + int32(16)
									return v206
								}
							}
						}
					}
				} else {
					v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v112 = F_BufFileSeek(m, v108, int32(0), int64(-4), int32(1))
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return int32(0)
					} else {
						if v112 != 0 {
							v114 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v114)
							v206 = v4
							m.G0 = v9 + int32(16)
							return v206
						} else {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v121 = F_BufFileReadMaybeEOF(m, v116, v9+int32(8), int32(4), int32(0))
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
								return int32(0)
							} else {
								v123 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								if v121 != 0 {
									v125 = v123
								} else {
									v125 = int32(0)
								}
								v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
								if v126 == int32(1) {
									v129 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v129)
									v175 = v125
									v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v178 = int32(0)
									v183 = F_BufFileSeek(m, v177, v178, base.I64_extend_i32_s(v178-v175), int32(1))
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return int32(0)
									} else {
										if v183 != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v230 = m.ExcPending
											if v230 != 0 {
												return int32(0)
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v232 = m.ExcPending
												if v232 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(381221), int32(0))
													mBase = m.M
													v236 = m.ExcPending
													if v236 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(491070), int32(1106), int32(377494))
														mBase = m.M
														v241 = m.ExcPending
														if v241 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v186 = m.T0[v185].(func(*base.Module, int32, int32) int32)(m, l0, v175)
											mBase = m.M
											v187 = m.ExcPending
											if v187 != 0 {
												return int32(0)
											} else {
												v206 = v186
												m.G0 = v9 + int32(16)
												return v206
											}
										}
									}
								} else {
									v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v137 = F_BufFileSeek(m, v131, int32(0), base.I64_extend_i32_s(int32(-8)-v125), int32(1))
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return int32(0)
									} else {
										if v137 != 0 {
											v139 = int32(0)
											v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v146 = F_BufFileSeek(m, v140, v139, base.I64_extend_i32_s(int32(-4)-v125), int32(1))
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return int32(0)
											} else {
												if v146 == int32(0) {
													v206 = v139
													m.G0 = v9 + int32(16)
													return v206
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int32(0)
													} else {
														F_errcode_for_file_access(m)
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(381221), int32(0))
															mBase = m.M
															v159 = m.ExcPending
															if v159 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(491070), int32(1089), int32(377494))
																mBase = m.M
																v164 = m.ExcPending
																if v164 != 0 {
																	return int32(0)
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
											v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v170 = F_BufFileReadMaybeEOF(m, v165, v9+int32(12), int32(4), int32(0))
											mBase = m.M
											v171 = m.ExcPending
											if v171 != 0 {
												return int32(0)
											} else {
												v172 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												if v170 != 0 {
													v174 = v172
												} else {
													v174 = int32(0)
												}
												v175 = v174
												v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v178 = int32(0)
												v183 = F_BufFileSeek(m, v177, v178, base.I64_extend_i32_s(v178-v175), int32(1))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return int32(0)
												} else {
													if v183 != 0 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v230 = m.ExcPending
														if v230 != 0 {
															return int32(0)
														} else {
															F_errcode_for_file_access(m)
															mBase = m.M
															v232 = m.ExcPending
															if v232 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(381221), int32(0))
																mBase = m.M
																v236 = m.ExcPending
																if v236 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(491070), int32(1106), int32(377494))
																	mBase = m.M
																	v241 = m.ExcPending
																	if v241 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
														v186 = m.T0[v185].(func(*base.Module, int32, int32) int32)(m, l0, v175)
														mBase = m.M
														v187 = m.ExcPending
														if v187 != 0 {
															return int32(0)
														} else {
															v206 = v186
															m.G0 = v9 + int32(16)
															return v206
														}
													}
												}
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
			v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
			if v58&int32(1) == int32(0) {
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
				*(*int32)(unsafe.Add(mBase, uint32(l0+int32(108)))) = v68
				v70 = *(*int64)(unsafe.Add(mBase, uint32(v63)+32))
				v71 = int64(*(*int32)(unsafe.Add(mBase, uint32(v63)+40)))
				*(*int64)(unsafe.Add(mBase, uint32(l0+int32(112)))) = v70 + v71
				v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
				if v74 == int32(0) {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
					v79 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
					v81 = F_BufFileSeek(m, v77, v78, v79, int32(0))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						if v81 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v215 = m.ExcPending
							if v215 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v217 = m.ExcPending
								if v217 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(381221), int32(0))
									mBase = m.M
									v221 = m.ExcPending
									if v221 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(491070), int32(1025), int32(377494))
										mBase = m.M
										v226 = m.ExcPending
										if v226 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2)
							v87 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v87)
							if l1 != 0 {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v90 = int32(4)
								v94 = F_BufFileReadMaybeEOF(m, v89, v9+v90, v90, int32(1))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									if v94 == int32(0) {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
										v206 = int32(0)
										m.G0 = v9 + int32(16)
										return v206
									} else {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
										if v98 == int32(0) {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
											v206 = int32(0)
											m.G0 = v9 + int32(16)
											return v206
										} else {
											v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v102 = m.T0[v101].(func(*base.Module, int32, int32) int32)(m, l0, v98)
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												v206 = v102
												m.G0 = v9 + int32(16)
												return v206
											}
										}
									}
								}
							} else {
								v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v112 = F_BufFileSeek(m, v108, int32(0), int64(-4), int32(1))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									if v112 != 0 {
										v114 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v114)
										v206 = v4
										m.G0 = v9 + int32(16)
										return v206
									} else {
										v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v121 = F_BufFileReadMaybeEOF(m, v116, v9+int32(8), int32(4), int32(0))
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int32(0)
										} else {
											v123 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
											if v121 != 0 {
												v125 = v123
											} else {
												v125 = int32(0)
											}
											v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
											if v126 == int32(1) {
												v129 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v129)
												v175 = v125
												v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v178 = int32(0)
												v183 = F_BufFileSeek(m, v177, v178, base.I64_extend_i32_s(v178-v175), int32(1))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return int32(0)
												} else {
													if v183 != 0 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v230 = m.ExcPending
														if v230 != 0 {
															return int32(0)
														} else {
															F_errcode_for_file_access(m)
															mBase = m.M
															v232 = m.ExcPending
															if v232 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(381221), int32(0))
																mBase = m.M
																v236 = m.ExcPending
																if v236 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(491070), int32(1106), int32(377494))
																	mBase = m.M
																	v241 = m.ExcPending
																	if v241 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
														v186 = m.T0[v185].(func(*base.Module, int32, int32) int32)(m, l0, v175)
														mBase = m.M
														v187 = m.ExcPending
														if v187 != 0 {
															return int32(0)
														} else {
															v206 = v186
															m.G0 = v9 + int32(16)
															return v206
														}
													}
												}
											} else {
												v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v137 = F_BufFileSeek(m, v131, int32(0), base.I64_extend_i32_s(int32(-8)-v125), int32(1))
												mBase = m.M
												v138 = m.ExcPending
												if v138 != 0 {
													return int32(0)
												} else {
													if v137 != 0 {
														v139 = int32(0)
														v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
														v146 = F_BufFileSeek(m, v140, v139, base.I64_extend_i32_s(int32(-4)-v125), int32(1))
														mBase = m.M
														v147 = m.ExcPending
														if v147 != 0 {
															return int32(0)
														} else {
															if v146 == int32(0) {
																v206 = v139
																m.G0 = v9 + int32(16)
																return v206
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v153 = m.ExcPending
																if v153 != 0 {
																	return int32(0)
																} else {
																	F_errcode_for_file_access(m)
																	mBase = m.M
																	v155 = m.ExcPending
																	if v155 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(381221), int32(0))
																		mBase = m.M
																		v159 = m.ExcPending
																		if v159 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(491070), int32(1089), int32(377494))
																			mBase = m.M
																			v164 = m.ExcPending
																			if v164 != 0 {
																				return int32(0)
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
														v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
														v170 = F_BufFileReadMaybeEOF(m, v165, v9+int32(12), int32(4), int32(0))
														mBase = m.M
														v171 = m.ExcPending
														if v171 != 0 {
															return int32(0)
														} else {
															v172 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
															if v170 != 0 {
																v174 = v172
															} else {
																v174 = int32(0)
															}
															v175 = v174
															v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
															v178 = int32(0)
															v183 = F_BufFileSeek(m, v177, v178, base.I64_extend_i32_s(v178-v175), int32(1))
															mBase = m.M
															v184 = m.ExcPending
															if v184 != 0 {
																return int32(0)
															} else {
																if v183 != 0 {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v230 = m.ExcPending
																	if v230 != 0 {
																		return int32(0)
																	} else {
																		F_errcode_for_file_access(m)
																		mBase = m.M
																		v232 = m.ExcPending
																		if v232 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(381221), int32(0))
																			mBase = m.M
																			v236 = m.ExcPending
																			if v236 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(491070), int32(1106), int32(377494))
																				mBase = m.M
																				v241 = m.ExcPending
																				if v241 != 0 {
																					return int32(0)
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			}
																		}
																	}
																} else {
																	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
																	v186 = m.T0[v185].(func(*base.Module, int32, int32) int32)(m, l0, v175)
																	mBase = m.M
																	v187 = m.ExcPending
																	if v187 != 0 {
																		return int32(0)
																	} else {
																		v206 = v186
																		m.G0 = v9 + int32(16)
																		return v206
																	}
																}
															}
														}
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
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2)
					v87 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v87)
					if l1 != 0 {
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v90 = int32(4)
						v94 = F_BufFileReadMaybeEOF(m, v89, v9+v90, v90, int32(1))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							if v94 == int32(0) {
								v105 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
								v206 = int32(0)
								m.G0 = v9 + int32(16)
								return v206
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
								if v98 == int32(0) {
									v105 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
									v206 = int32(0)
									m.G0 = v9 + int32(16)
									return v206
								} else {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
									v102 = m.T0[v101].(func(*base.Module, int32, int32) int32)(m, l0, v98)
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return int32(0)
									} else {
										v206 = v102
										m.G0 = v9 + int32(16)
										return v206
									}
								}
							}
						}
					} else {
						v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v112 = F_BufFileSeek(m, v108, int32(0), int64(-4), int32(1))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							if v112 != 0 {
								v114 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v114)
								v206 = v4
								m.G0 = v9 + int32(16)
								return v206
							} else {
								v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v121 = F_BufFileReadMaybeEOF(m, v116, v9+int32(8), int32(4), int32(0))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int32(0)
								} else {
									v123 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
									if v121 != 0 {
										v125 = v123
									} else {
										v125 = int32(0)
									}
									v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
									if v126 == int32(1) {
										v129 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v129)
										v175 = v125
										v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v178 = int32(0)
										v183 = F_BufFileSeek(m, v177, v178, base.I64_extend_i32_s(v178-v175), int32(1))
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return int32(0)
										} else {
											if v183 != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v230 = m.ExcPending
												if v230 != 0 {
													return int32(0)
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v232 = m.ExcPending
													if v232 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(381221), int32(0))
														mBase = m.M
														v236 = m.ExcPending
														if v236 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(491070), int32(1106), int32(377494))
															mBase = m.M
															v241 = m.ExcPending
															if v241 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v186 = m.T0[v185].(func(*base.Module, int32, int32) int32)(m, l0, v175)
												mBase = m.M
												v187 = m.ExcPending
												if v187 != 0 {
													return int32(0)
												} else {
													v206 = v186
													m.G0 = v9 + int32(16)
													return v206
												}
											}
										}
									} else {
										v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v137 = F_BufFileSeek(m, v131, int32(0), base.I64_extend_i32_s(int32(-8)-v125), int32(1))
										mBase = m.M
										v138 = m.ExcPending
										if v138 != 0 {
											return int32(0)
										} else {
											if v137 != 0 {
												v139 = int32(0)
												v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v146 = F_BufFileSeek(m, v140, v139, base.I64_extend_i32_s(int32(-4)-v125), int32(1))
												mBase = m.M
												v147 = m.ExcPending
												if v147 != 0 {
													return int32(0)
												} else {
													if v146 == int32(0) {
														v206 = v139
														m.G0 = v9 + int32(16)
														return v206
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int32(0)
														} else {
															F_errcode_for_file_access(m)
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(381221), int32(0))
																mBase = m.M
																v159 = m.ExcPending
																if v159 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(491070), int32(1089), int32(377494))
																	mBase = m.M
																	v164 = m.ExcPending
																	if v164 != 0 {
																		return int32(0)
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
												v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v170 = F_BufFileReadMaybeEOF(m, v165, v9+int32(12), int32(4), int32(0))
												mBase = m.M
												v171 = m.ExcPending
												if v171 != 0 {
													return int32(0)
												} else {
													v172 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
													if v170 != 0 {
														v174 = v172
													} else {
														v174 = int32(0)
													}
													v175 = v174
													v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													v178 = int32(0)
													v183 = F_BufFileSeek(m, v177, v178, base.I64_extend_i32_s(v178-v175), int32(1))
													mBase = m.M
													v184 = m.ExcPending
													if v184 != 0 {
														return int32(0)
													} else {
														if v183 != 0 {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v230 = m.ExcPending
															if v230 != 0 {
																return int32(0)
															} else {
																F_errcode_for_file_access(m)
																mBase = m.M
																v232 = m.ExcPending
																if v232 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(381221), int32(0))
																	mBase = m.M
																	v236 = m.ExcPending
																	if v236 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(491070), int32(1106), int32(377494))
																		mBase = m.M
																		v241 = m.ExcPending
																		if v241 != 0 {
																			return int32(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																}
															}
														} else {
															v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
															v186 = m.T0[v185].(func(*base.Module, int32, int32) int32)(m, l0, v175)
															mBase = m.M
															v187 = m.ExcPending
															if v187 != 0 {
																return int32(0)
															} else {
																v206 = v186
																m.G0 = v9 + int32(16)
																return v206
															}
														}
													}
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
				v206 = v4
				m.G0 = v9 + int32(16)
				return v206
			}
		}
	case 2:
		v87 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v87)
		if l1 != 0 {
			v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v90 = int32(4)
			v94 = F_BufFileReadMaybeEOF(m, v89, v9+v90, v90, int32(1))
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return int32(0)
			} else {
				if v94 == int32(0) {
					v105 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
					v206 = int32(0)
					m.G0 = v9 + int32(16)
					return v206
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					if v98 == int32(0) {
						v105 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
						v206 = int32(0)
						m.G0 = v9 + int32(16)
						return v206
					} else {
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
						v102 = m.T0[v101].(func(*base.Module, int32, int32) int32)(m, l0, v98)
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return int32(0)
						} else {
							v206 = v102
							m.G0 = v9 + int32(16)
							return v206
						}
					}
				}
			}
		} else {
			v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v112 = F_BufFileSeek(m, v108, int32(0), int64(-4), int32(1))
			mBase = m.M
			v113 = m.ExcPending
			if v113 != 0 {
				return int32(0)
			} else {
				if v112 != 0 {
					v114 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v114)
					v206 = v4
					m.G0 = v9 + int32(16)
					return v206
				} else {
					v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v121 = F_BufFileReadMaybeEOF(m, v116, v9+int32(8), int32(4), int32(0))
					mBase = m.M
					v122 = m.ExcPending
					if v122 != 0 {
						return int32(0)
					} else {
						v123 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						if v121 != 0 {
							v125 = v123
						} else {
							v125 = int32(0)
						}
						v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
						if v126 == int32(1) {
							v129 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v129)
							v175 = v125
							v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v178 = int32(0)
							v183 = F_BufFileSeek(m, v177, v178, base.I64_extend_i32_s(v178-v175), int32(1))
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return int32(0)
							} else {
								if v183 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v230 = m.ExcPending
									if v230 != 0 {
										return int32(0)
									} else {
										F_errcode_for_file_access(m)
										mBase = m.M
										v232 = m.ExcPending
										if v232 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(381221), int32(0))
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(491070), int32(1106), int32(377494))
												mBase = m.M
												v241 = m.ExcPending
												if v241 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
									v186 = m.T0[v185].(func(*base.Module, int32, int32) int32)(m, l0, v175)
									mBase = m.M
									v187 = m.ExcPending
									if v187 != 0 {
										return int32(0)
									} else {
										v206 = v186
										m.G0 = v9 + int32(16)
										return v206
									}
								}
							}
						} else {
							v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v137 = F_BufFileSeek(m, v131, int32(0), base.I64_extend_i32_s(int32(-8)-v125), int32(1))
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return int32(0)
							} else {
								if v137 != 0 {
									v139 = int32(0)
									v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v146 = F_BufFileSeek(m, v140, v139, base.I64_extend_i32_s(int32(-4)-v125), int32(1))
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return int32(0)
									} else {
										if v146 == int32(0) {
											v206 = v139
											m.G0 = v9 + int32(16)
											return v206
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v153 = m.ExcPending
											if v153 != 0 {
												return int32(0)
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(381221), int32(0))
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(491070), int32(1089), int32(377494))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return int32(0)
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
									v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v170 = F_BufFileReadMaybeEOF(m, v165, v9+int32(12), int32(4), int32(0))
									mBase = m.M
									v171 = m.ExcPending
									if v171 != 0 {
										return int32(0)
									} else {
										v172 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
										if v170 != 0 {
											v174 = v172
										} else {
											v174 = int32(0)
										}
										v175 = v174
										v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v178 = int32(0)
										v183 = F_BufFileSeek(m, v177, v178, base.I64_extend_i32_s(v178-v175), int32(1))
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return int32(0)
										} else {
											if v183 != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v230 = m.ExcPending
												if v230 != 0 {
													return int32(0)
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v232 = m.ExcPending
													if v232 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(381221), int32(0))
														mBase = m.M
														v236 = m.ExcPending
														if v236 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(491070), int32(1106), int32(377494))
															mBase = m.M
															v241 = m.ExcPending
															if v241 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v186 = m.T0[v185].(func(*base.Module, int32, int32) int32)(m, l0, v175)
												mBase = m.M
												v187 = m.ExcPending
												if v187 != 0 {
													return int32(0)
												} else {
													v206 = v186
													m.G0 = v9 + int32(16)
													return v206
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v191 = m.ExcPending
		if v191 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(347462), int32(0))
			mBase = m.M
			v195 = m.ExcPending
			if v195 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(491070), int32(1111), int32(377494))
				mBase = m.M
				v200 = m.ExcPending
				if v200 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
