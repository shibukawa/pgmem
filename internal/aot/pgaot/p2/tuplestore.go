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
	var v22 int32
	_ = v22
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
			if v16&int32(1) == int32(0) {
				m.G0 = v6 + int32(16)
				return base.B2i32(v10 != int32(0))
			} else {
				F_pfree(m, v10)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
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
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v37
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v39
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
				*(*int64)(unsafe.Add(mBase, uint32(v36))) = v41
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
					v37 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v37
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v39
					v41 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
					*(*int64)(unsafe.Add(mBase, uint32(v36))) = v41
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
					F_errmsg_internal(m, int32(_a_F_tuplestore_alloc_read_pointer_0), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_tuplestore_alloc_read_pointer_1), int32(401), int32(_a_F_tuplestore_alloc_read_pointer_2))
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
					v37 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v37
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v39
					v41 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
					*(*int64)(unsafe.Add(mBase, uint32(v36))) = v41
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
						v37 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v37
						v39 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v39
						v41 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
						*(*int64)(unsafe.Add(mBase, uint32(v36))) = v41
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
				F_errmsg_internal(m, int32(_a_F_tuplestore_alloc_read_pointer_0), int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_tuplestore_alloc_read_pointer_1), int32(401), int32(_a_F_tuplestore_alloc_read_pointer_2))
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
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v37
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v39
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
				*(*int64)(unsafe.Add(mBase, uint32(v36))) = v41
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
					v37 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v37
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v39
					v41 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
					*(*int64)(unsafe.Add(mBase, uint32(v36))) = v41
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
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
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int64
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	if l1 == l2 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
		v14 = int32(24)
		v16 = v13 + l1*v14
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		v20 = v13 + l2*v14
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v21
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v23
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		v26 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
		*(*int64)(unsafe.Add(mBase, uint32(v20))) = v26
		if v25 != v17 {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			if v31 < int32(2) {
				v114 = v30
			} else {
				v35 = v31 - int32(1)
				v36 = int32(3)
				v37 = v35 & v36
				if base.Ui32(v31-int32(2)) < base.Ui32(v36) {
					v78 = int32(1)
					v81 = v30
					v90 = v78
					v92 = int32(0)
					v93 = v81
					for {
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v29+v90*int32(24))))
						v102 = v101 | v93
						v103 = int32(1)
						v106 = v92 + v103
						if v106 != v37 {
							v90 = v90 + v103
							v92 = v106
							v93 = v102
							continue
						} else {
							break
						}
						break
					}
					v114 = v102
				} else {
					v49 = int32(1)
					v52 = v30
					v56 = int32(0)
					for {
						v59 = v29 + v49*int32(24)
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+72))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+48))
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
						v67 = v60 | (v61 | (v62 | (v63 | v52)))
						v68 = int32(4)
						v69 = v49 + v68
						v71 = v56 + v68
						if v71 != v35&int32(-4) {
							v49 = v69
							v52 = v67
							v56 = v71
							continue
						} else {
							break
						}
						break
					}
					if v37 == int32(0) {
						v114 = v67
					} else {
						v78 = v69
						v81 = v67
						v90 = v78
						v92 = int32(0)
						v93 = v81
						for {
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v29+v90*int32(24))))
							v102 = v101 | v93
							v103 = int32(1)
							v106 = v92 + v103
							if v106 != v37 {
								v90 = v90 + v103
								v92 = v106
								v93 = v102
								continue
							} else {
								break
							}
							break
						}
						v114 = v102
					}
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
		} else {
		}
		v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(v131) < base.Ui32(int32(2)) {
			return
		} else {
			if v131 != int32(2) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v213 = m.ExcPending
				if v213 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_tuplestore_copy_read_pointer_0), int32(0))
					mBase = m.M
					v217 = m.ExcPending
					if v217 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_tuplestore_copy_read_pointer_1), int32(1394), int32(_a_F_tuplestore_copy_read_pointer_2))
						mBase = m.M
						v222 = m.ExcPending
						if v222 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				if l2 == v136 {
					v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
					if v139 == int32(1) {
						v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
						v143 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
						v145 = F_BufFileSeek(m, v138, v142, v143, int32(0))
						mBase = m.M
						v146 = m.ExcPending
						if v146 != 0 {
							return
						} else {
							if v145 == int32(0) {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v154 = m.ExcPending
									if v154 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_tuplestore_copy_read_pointer_3), int32(0))
										mBase = m.M
										v158 = m.ExcPending
										if v158 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_tuplestore_copy_read_pointer_1), int32(1373), int32(_a_F_tuplestore_copy_read_pointer_2))
											mBase = m.M
											v163 = m.ExcPending
											if v163 != 0 {
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
						v164 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
						v165 = *(*int64)(unsafe.Add(mBase, uint32(v20)+16))
						v167 = F_BufFileSeek(m, v138, v164, v165, int32(0))
						mBase = m.M
						v168 = m.ExcPending
						if v168 != 0 {
							return
						} else {
							if v167 == int32(0) {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v174 = m.ExcPending
								if v174 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v176 = m.ExcPending
									if v176 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_tuplestore_copy_read_pointer_3), int32(0))
										mBase = m.M
										v180 = m.ExcPending
										if v180 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_tuplestore_copy_read_pointer_1), int32(1382), int32(_a_F_tuplestore_copy_read_pointer_2))
											mBase = m.M
											v185 = m.ExcPending
											if v185 != 0 {
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
					if l1 != v136 {
					} else {
						v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
						if v187 != 0 {
						} else {
							v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v193 = *(*int32)(unsafe.Add(mBase, uint32(v188)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v20+int32(12)))) = v193
							v195 = *(*int64)(unsafe.Add(mBase, uint32(v188)+32))
							v196 = int64(*(*int32)(unsafe.Add(mBase, uint32(v188)+40)))
							*(*int64)(unsafe.Add(mBase, uint32(v20+int32(16)))) = v195 + v196
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
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
		if v18 != 0 {
			v19 = int32(_a_F_tuplestore_get_stats_0)
		} else {
			v19 = int32(_a_F_tuplestore_get_stats_1)
		}
		v35 = v19
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
		v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int64)(unsafe.Add(mBase, uint32(l2))) = v37
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v21 = F_BufFileSize(m, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			if v21 < v6 {
				v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
				v28 = v24
				v29 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v29)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v28
				v35 = int32(_a_F_tuplestore_get_stats_0)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
				v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int64)(unsafe.Add(mBase, uint32(l2))) = v37
				return
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v26 = F_BufFileSize(m, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v28 = v26
					v29 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v29)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v28
					v35 = int32(_a_F_tuplestore_get_stats_0)
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
					v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = v37
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
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
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
				v202 = v4
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
				if v23 <= v22 {
					v198 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v198)
					v202 = v4
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v22 + int32(1)
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v25+v22<<(uint(int32(2))%32))))
					v202 = v32
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
					v202 = int32(0)
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+v44<<(uint(int32(2))%32)-int32(4))))
					v202 = v55
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				if v38 <= v39 {
					v202 = v4
				} else {
					v44 = v38 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v44
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					if v44 <= v46 {
						v202 = int32(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+v44<<(uint(int32(2))%32)-int32(4))))
						v202 = v55
					}
				}
			}
		}
		m.G0 = v9 + int32(16)
		return v202
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
						v210 = m.ExcPending
						if v210 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v212 = m.ExcPending
							if v212 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
								mBase = m.M
								v216 = m.ExcPending
								if v216 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1025), int32(_a_F_tuplestore_gettuple_2))
									mBase = m.M
									v221 = m.ExcPending
									if v221 != 0 {
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
							v94 = F_BufFileReadCommon(m, v89, v9+v90, v90, int32(1))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								if v94 == int32(0) {
									v105 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
									v202 = v4
									m.G0 = v9 + int32(16)
									return v202
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
									if v98 == int32(0) {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
										v202 = v4
										m.G0 = v9 + int32(16)
										return v202
									} else {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
										v102 = m.T0[v101].(func(*base.Module, int32, int32) int32)(m, l0, v98)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int32(0)
										} else {
											v202 = v102
											m.G0 = v9 + int32(16)
											return v202
										}
									}
								}
							}
						} else {
							v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v111 = F_BufFileSeek(m, v107, int32(0), int64(-4), int32(1))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								if v111 != 0 {
									v113 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v113)
									v202 = v4
									m.G0 = v9 + int32(16)
									return v202
								} else {
									v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v120 = F_BufFileReadCommon(m, v115, v9+int32(8), int32(4), int32(0))
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										v122 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
										if v120 != 0 {
											v124 = v122
										} else {
											v124 = int32(0)
										}
										v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
										if v125 == int32(1) {
											v128 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v128)
											v173 = v124
											v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v175 = int32(0)
											v180 = F_BufFileSeek(m, v174, v175, base.I64_extend_i32_s(v175-v173), int32(1))
											mBase = m.M
											v181 = m.ExcPending
											if v181 != 0 {
												return int32(0)
											} else {
												if v180 != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v225 = m.ExcPending
													if v225 != 0 {
														return int32(0)
													} else {
														F_errcode_for_file_access(m)
														mBase = m.M
														v227 = m.ExcPending
														if v227 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
															mBase = m.M
															v231 = m.ExcPending
															if v231 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1106), int32(_a_F_tuplestore_gettuple_2))
																mBase = m.M
																v236 = m.ExcPending
																if v236 != 0 {
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
													v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
													v183 = m.T0[v182].(func(*base.Module, int32, int32) int32)(m, l0, v173)
													mBase = m.M
													v184 = m.ExcPending
													if v184 != 0 {
														return int32(0)
													} else {
														v202 = v183
														m.G0 = v9 + int32(16)
														return v202
													}
												}
											}
										} else {
											v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v136 = F_BufFileSeek(m, v130, int32(0), base.I64_extend_i32_s(int32(-8)-v124), int32(1))
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return int32(0)
											} else {
												if v136 != 0 {
													v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													v144 = F_BufFileSeek(m, v138, int32(0), base.I64_extend_i32_s(int32(-4)-v124), int32(1))
													mBase = m.M
													v145 = m.ExcPending
													if v145 != 0 {
														return int32(0)
													} else {
														if v144 == int32(0) {
															v202 = v4
															m.G0 = v9 + int32(16)
															return v202
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v151 = m.ExcPending
															if v151 != 0 {
																return int32(0)
															} else {
																F_errcode_for_file_access(m)
																mBase = m.M
																v153 = m.ExcPending
																if v153 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
																	mBase = m.M
																	v157 = m.ExcPending
																	if v157 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1089), int32(_a_F_tuplestore_gettuple_2))
																		mBase = m.M
																		v162 = m.ExcPending
																		if v162 != 0 {
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
													v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													v168 = F_BufFileReadCommon(m, v163, v9+int32(12), int32(4), int32(0))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return int32(0)
													} else {
														v170 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
														if v168 != 0 {
															v172 = v170
														} else {
															v172 = int32(0)
														}
														v173 = v172
														v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
														v175 = int32(0)
														v180 = F_BufFileSeek(m, v174, v175, base.I64_extend_i32_s(v175-v173), int32(1))
														mBase = m.M
														v181 = m.ExcPending
														if v181 != 0 {
															return int32(0)
														} else {
															if v180 != 0 {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v225 = m.ExcPending
																if v225 != 0 {
																	return int32(0)
																} else {
																	F_errcode_for_file_access(m)
																	mBase = m.M
																	v227 = m.ExcPending
																	if v227 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
																		mBase = m.M
																		v231 = m.ExcPending
																		if v231 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1106), int32(_a_F_tuplestore_gettuple_2))
																			mBase = m.M
																			v236 = m.ExcPending
																			if v236 != 0 {
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
																v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
																v183 = m.T0[v182].(func(*base.Module, int32, int32) int32)(m, l0, v173)
																mBase = m.M
																v184 = m.ExcPending
																if v184 != 0 {
																	return int32(0)
																} else {
																	v202 = v183
																	m.G0 = v9 + int32(16)
																	return v202
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
					v94 = F_BufFileReadCommon(m, v89, v9+v90, v90, int32(1))
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						if v94 == int32(0) {
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
							v202 = v4
							m.G0 = v9 + int32(16)
							return v202
						} else {
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							if v98 == int32(0) {
								v105 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
								v202 = v4
								m.G0 = v9 + int32(16)
								return v202
							} else {
								v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
								v102 = m.T0[v101].(func(*base.Module, int32, int32) int32)(m, l0, v98)
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int32(0)
								} else {
									v202 = v102
									m.G0 = v9 + int32(16)
									return v202
								}
							}
						}
					}
				} else {
					v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v111 = F_BufFileSeek(m, v107, int32(0), int64(-4), int32(1))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return int32(0)
					} else {
						if v111 != 0 {
							v113 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v113)
							v202 = v4
							m.G0 = v9 + int32(16)
							return v202
						} else {
							v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v120 = F_BufFileReadCommon(m, v115, v9+int32(8), int32(4), int32(0))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								if v120 != 0 {
									v124 = v122
								} else {
									v124 = int32(0)
								}
								v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
								if v125 == int32(1) {
									v128 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v128)
									v173 = v124
									v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v175 = int32(0)
									v180 = F_BufFileSeek(m, v174, v175, base.I64_extend_i32_s(v175-v173), int32(1))
									mBase = m.M
									v181 = m.ExcPending
									if v181 != 0 {
										return int32(0)
									} else {
										if v180 != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v225 = m.ExcPending
											if v225 != 0 {
												return int32(0)
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v227 = m.ExcPending
												if v227 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
													mBase = m.M
													v231 = m.ExcPending
													if v231 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1106), int32(_a_F_tuplestore_gettuple_2))
														mBase = m.M
														v236 = m.ExcPending
														if v236 != 0 {
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
											v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v183 = m.T0[v182].(func(*base.Module, int32, int32) int32)(m, l0, v173)
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return int32(0)
											} else {
												v202 = v183
												m.G0 = v9 + int32(16)
												return v202
											}
										}
									}
								} else {
									v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v136 = F_BufFileSeek(m, v130, int32(0), base.I64_extend_i32_s(int32(-8)-v124), int32(1))
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return int32(0)
									} else {
										if v136 != 0 {
											v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v144 = F_BufFileSeek(m, v138, int32(0), base.I64_extend_i32_s(int32(-4)-v124), int32(1))
											mBase = m.M
											v145 = m.ExcPending
											if v145 != 0 {
												return int32(0)
											} else {
												if v144 == int32(0) {
													v202 = v4
													m.G0 = v9 + int32(16)
													return v202
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v151 = m.ExcPending
													if v151 != 0 {
														return int32(0)
													} else {
														F_errcode_for_file_access(m)
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1089), int32(_a_F_tuplestore_gettuple_2))
																mBase = m.M
																v162 = m.ExcPending
																if v162 != 0 {
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
											v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v168 = F_BufFileReadCommon(m, v163, v9+int32(12), int32(4), int32(0))
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
												return int32(0)
											} else {
												v170 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												if v168 != 0 {
													v172 = v170
												} else {
													v172 = int32(0)
												}
												v173 = v172
												v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v175 = int32(0)
												v180 = F_BufFileSeek(m, v174, v175, base.I64_extend_i32_s(v175-v173), int32(1))
												mBase = m.M
												v181 = m.ExcPending
												if v181 != 0 {
													return int32(0)
												} else {
													if v180 != 0 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v225 = m.ExcPending
														if v225 != 0 {
															return int32(0)
														} else {
															F_errcode_for_file_access(m)
															mBase = m.M
															v227 = m.ExcPending
															if v227 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
																mBase = m.M
																v231 = m.ExcPending
																if v231 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1106), int32(_a_F_tuplestore_gettuple_2))
																	mBase = m.M
																	v236 = m.ExcPending
																	if v236 != 0 {
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
														v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
														v183 = m.T0[v182].(func(*base.Module, int32, int32) int32)(m, l0, v173)
														mBase = m.M
														v184 = m.ExcPending
														if v184 != 0 {
															return int32(0)
														} else {
															v202 = v183
															m.G0 = v9 + int32(16)
															return v202
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
							v210 = m.ExcPending
							if v210 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v212 = m.ExcPending
								if v212 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
									mBase = m.M
									v216 = m.ExcPending
									if v216 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1025), int32(_a_F_tuplestore_gettuple_2))
										mBase = m.M
										v221 = m.ExcPending
										if v221 != 0 {
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
								v94 = F_BufFileReadCommon(m, v89, v9+v90, v90, int32(1))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									if v94 == int32(0) {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
										v202 = v4
										m.G0 = v9 + int32(16)
										return v202
									} else {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
										if v98 == int32(0) {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
											v202 = v4
											m.G0 = v9 + int32(16)
											return v202
										} else {
											v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v102 = m.T0[v101].(func(*base.Module, int32, int32) int32)(m, l0, v98)
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												v202 = v102
												m.G0 = v9 + int32(16)
												return v202
											}
										}
									}
								}
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v111 = F_BufFileSeek(m, v107, int32(0), int64(-4), int32(1))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									if v111 != 0 {
										v113 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v113)
										v202 = v4
										m.G0 = v9 + int32(16)
										return v202
									} else {
										v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v120 = F_BufFileReadCommon(m, v115, v9+int32(8), int32(4), int32(0))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											v122 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
											if v120 != 0 {
												v124 = v122
											} else {
												v124 = int32(0)
											}
											v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
											if v125 == int32(1) {
												v128 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v128)
												v173 = v124
												v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v175 = int32(0)
												v180 = F_BufFileSeek(m, v174, v175, base.I64_extend_i32_s(v175-v173), int32(1))
												mBase = m.M
												v181 = m.ExcPending
												if v181 != 0 {
													return int32(0)
												} else {
													if v180 != 0 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v225 = m.ExcPending
														if v225 != 0 {
															return int32(0)
														} else {
															F_errcode_for_file_access(m)
															mBase = m.M
															v227 = m.ExcPending
															if v227 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
																mBase = m.M
																v231 = m.ExcPending
																if v231 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1106), int32(_a_F_tuplestore_gettuple_2))
																	mBase = m.M
																	v236 = m.ExcPending
																	if v236 != 0 {
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
														v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
														v183 = m.T0[v182].(func(*base.Module, int32, int32) int32)(m, l0, v173)
														mBase = m.M
														v184 = m.ExcPending
														if v184 != 0 {
															return int32(0)
														} else {
															v202 = v183
															m.G0 = v9 + int32(16)
															return v202
														}
													}
												}
											} else {
												v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v136 = F_BufFileSeek(m, v130, int32(0), base.I64_extend_i32_s(int32(-8)-v124), int32(1))
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return int32(0)
												} else {
													if v136 != 0 {
														v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
														v144 = F_BufFileSeek(m, v138, int32(0), base.I64_extend_i32_s(int32(-4)-v124), int32(1))
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
															return int32(0)
														} else {
															if v144 == int32(0) {
																v202 = v4
																m.G0 = v9 + int32(16)
																return v202
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v151 = m.ExcPending
																if v151 != 0 {
																	return int32(0)
																} else {
																	F_errcode_for_file_access(m)
																	mBase = m.M
																	v153 = m.ExcPending
																	if v153 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
																		mBase = m.M
																		v157 = m.ExcPending
																		if v157 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1089), int32(_a_F_tuplestore_gettuple_2))
																			mBase = m.M
																			v162 = m.ExcPending
																			if v162 != 0 {
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
														v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
														v168 = F_BufFileReadCommon(m, v163, v9+int32(12), int32(4), int32(0))
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return int32(0)
														} else {
															v170 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
															if v168 != 0 {
																v172 = v170
															} else {
																v172 = int32(0)
															}
															v173 = v172
															v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
															v175 = int32(0)
															v180 = F_BufFileSeek(m, v174, v175, base.I64_extend_i32_s(v175-v173), int32(1))
															mBase = m.M
															v181 = m.ExcPending
															if v181 != 0 {
																return int32(0)
															} else {
																if v180 != 0 {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v225 = m.ExcPending
																	if v225 != 0 {
																		return int32(0)
																	} else {
																		F_errcode_for_file_access(m)
																		mBase = m.M
																		v227 = m.ExcPending
																		if v227 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
																			mBase = m.M
																			v231 = m.ExcPending
																			if v231 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1106), int32(_a_F_tuplestore_gettuple_2))
																				mBase = m.M
																				v236 = m.ExcPending
																				if v236 != 0 {
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
																	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
																	v183 = m.T0[v182].(func(*base.Module, int32, int32) int32)(m, l0, v173)
																	mBase = m.M
																	v184 = m.ExcPending
																	if v184 != 0 {
																		return int32(0)
																	} else {
																		v202 = v183
																		m.G0 = v9 + int32(16)
																		return v202
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
						v94 = F_BufFileReadCommon(m, v89, v9+v90, v90, int32(1))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							if v94 == int32(0) {
								v105 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
								v202 = v4
								m.G0 = v9 + int32(16)
								return v202
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
								if v98 == int32(0) {
									v105 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
									v202 = v4
									m.G0 = v9 + int32(16)
									return v202
								} else {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
									v102 = m.T0[v101].(func(*base.Module, int32, int32) int32)(m, l0, v98)
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return int32(0)
									} else {
										v202 = v102
										m.G0 = v9 + int32(16)
										return v202
									}
								}
							}
						}
					} else {
						v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v111 = F_BufFileSeek(m, v107, int32(0), int64(-4), int32(1))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							if v111 != 0 {
								v113 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v113)
								v202 = v4
								m.G0 = v9 + int32(16)
								return v202
							} else {
								v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v120 = F_BufFileReadCommon(m, v115, v9+int32(8), int32(4), int32(0))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									v122 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
									if v120 != 0 {
										v124 = v122
									} else {
										v124 = int32(0)
									}
									v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
									if v125 == int32(1) {
										v128 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v128)
										v173 = v124
										v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v175 = int32(0)
										v180 = F_BufFileSeek(m, v174, v175, base.I64_extend_i32_s(v175-v173), int32(1))
										mBase = m.M
										v181 = m.ExcPending
										if v181 != 0 {
											return int32(0)
										} else {
											if v180 != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v225 = m.ExcPending
												if v225 != 0 {
													return int32(0)
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v227 = m.ExcPending
													if v227 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
														mBase = m.M
														v231 = m.ExcPending
														if v231 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1106), int32(_a_F_tuplestore_gettuple_2))
															mBase = m.M
															v236 = m.ExcPending
															if v236 != 0 {
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
												v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v183 = m.T0[v182].(func(*base.Module, int32, int32) int32)(m, l0, v173)
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return int32(0)
												} else {
													v202 = v183
													m.G0 = v9 + int32(16)
													return v202
												}
											}
										}
									} else {
										v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v136 = F_BufFileSeek(m, v130, int32(0), base.I64_extend_i32_s(int32(-8)-v124), int32(1))
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return int32(0)
										} else {
											if v136 != 0 {
												v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v144 = F_BufFileSeek(m, v138, int32(0), base.I64_extend_i32_s(int32(-4)-v124), int32(1))
												mBase = m.M
												v145 = m.ExcPending
												if v145 != 0 {
													return int32(0)
												} else {
													if v144 == int32(0) {
														v202 = v4
														m.G0 = v9 + int32(16)
														return v202
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v151 = m.ExcPending
														if v151 != 0 {
															return int32(0)
														} else {
															F_errcode_for_file_access(m)
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1089), int32(_a_F_tuplestore_gettuple_2))
																	mBase = m.M
																	v162 = m.ExcPending
																	if v162 != 0 {
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
												v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v168 = F_BufFileReadCommon(m, v163, v9+int32(12), int32(4), int32(0))
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return int32(0)
												} else {
													v170 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
													if v168 != 0 {
														v172 = v170
													} else {
														v172 = int32(0)
													}
													v173 = v172
													v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													v175 = int32(0)
													v180 = F_BufFileSeek(m, v174, v175, base.I64_extend_i32_s(v175-v173), int32(1))
													mBase = m.M
													v181 = m.ExcPending
													if v181 != 0 {
														return int32(0)
													} else {
														if v180 != 0 {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v225 = m.ExcPending
															if v225 != 0 {
																return int32(0)
															} else {
																F_errcode_for_file_access(m)
																mBase = m.M
																v227 = m.ExcPending
																if v227 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
																	mBase = m.M
																	v231 = m.ExcPending
																	if v231 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1106), int32(_a_F_tuplestore_gettuple_2))
																		mBase = m.M
																		v236 = m.ExcPending
																		if v236 != 0 {
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
															v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
															v183 = m.T0[v182].(func(*base.Module, int32, int32) int32)(m, l0, v173)
															mBase = m.M
															v184 = m.ExcPending
															if v184 != 0 {
																return int32(0)
															} else {
																v202 = v183
																m.G0 = v9 + int32(16)
																return v202
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
				v202 = v4
				m.G0 = v9 + int32(16)
				return v202
			}
		}
	case 2:
		v87 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v87)
		if l1 != 0 {
			v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v90 = int32(4)
			v94 = F_BufFileReadCommon(m, v89, v9+v90, v90, int32(1))
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return int32(0)
			} else {
				if v94 == int32(0) {
					v105 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
					v202 = v4
					m.G0 = v9 + int32(16)
					return v202
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					if v98 == int32(0) {
						v105 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v105)
						v202 = v4
						m.G0 = v9 + int32(16)
						return v202
					} else {
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
						v102 = m.T0[v101].(func(*base.Module, int32, int32) int32)(m, l0, v98)
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return int32(0)
						} else {
							v202 = v102
							m.G0 = v9 + int32(16)
							return v202
						}
					}
				}
			}
		} else {
			v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v111 = F_BufFileSeek(m, v107, int32(0), int64(-4), int32(1))
			mBase = m.M
			v112 = m.ExcPending
			if v112 != 0 {
				return int32(0)
			} else {
				if v111 != 0 {
					v113 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v113)
					v202 = v4
					m.G0 = v9 + int32(16)
					return v202
				} else {
					v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v120 = F_BufFileReadCommon(m, v115, v9+int32(8), int32(4), int32(0))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						if v120 != 0 {
							v124 = v122
						} else {
							v124 = int32(0)
						}
						v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
						if v125 == int32(1) {
							v128 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v128)
							v173 = v124
							v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v175 = int32(0)
							v180 = F_BufFileSeek(m, v174, v175, base.I64_extend_i32_s(v175-v173), int32(1))
							mBase = m.M
							v181 = m.ExcPending
							if v181 != 0 {
								return int32(0)
							} else {
								if v180 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v225 = m.ExcPending
									if v225 != 0 {
										return int32(0)
									} else {
										F_errcode_for_file_access(m)
										mBase = m.M
										v227 = m.ExcPending
										if v227 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
											mBase = m.M
											v231 = m.ExcPending
											if v231 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1106), int32(_a_F_tuplestore_gettuple_2))
												mBase = m.M
												v236 = m.ExcPending
												if v236 != 0 {
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
									v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
									v183 = m.T0[v182].(func(*base.Module, int32, int32) int32)(m, l0, v173)
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return int32(0)
									} else {
										v202 = v183
										m.G0 = v9 + int32(16)
										return v202
									}
								}
							}
						} else {
							v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v136 = F_BufFileSeek(m, v130, int32(0), base.I64_extend_i32_s(int32(-8)-v124), int32(1))
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return int32(0)
							} else {
								if v136 != 0 {
									v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v144 = F_BufFileSeek(m, v138, int32(0), base.I64_extend_i32_s(int32(-4)-v124), int32(1))
									mBase = m.M
									v145 = m.ExcPending
									if v145 != 0 {
										return int32(0)
									} else {
										if v144 == int32(0) {
											v202 = v4
											m.G0 = v9 + int32(16)
											return v202
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v151 = m.ExcPending
											if v151 != 0 {
												return int32(0)
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v153 = m.ExcPending
												if v153 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1089), int32(_a_F_tuplestore_gettuple_2))
														mBase = m.M
														v162 = m.ExcPending
														if v162 != 0 {
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
									v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v168 = F_BufFileReadCommon(m, v163, v9+int32(12), int32(4), int32(0))
									mBase = m.M
									v169 = m.ExcPending
									if v169 != 0 {
										return int32(0)
									} else {
										v170 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
										if v168 != 0 {
											v172 = v170
										} else {
											v172 = int32(0)
										}
										v173 = v172
										v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v175 = int32(0)
										v180 = F_BufFileSeek(m, v174, v175, base.I64_extend_i32_s(v175-v173), int32(1))
										mBase = m.M
										v181 = m.ExcPending
										if v181 != 0 {
											return int32(0)
										} else {
											if v180 != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v225 = m.ExcPending
												if v225 != 0 {
													return int32(0)
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v227 = m.ExcPending
													if v227 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_tuplestore_gettuple_0), int32(0))
														mBase = m.M
														v231 = m.ExcPending
														if v231 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1106), int32(_a_F_tuplestore_gettuple_2))
															mBase = m.M
															v236 = m.ExcPending
															if v236 != 0 {
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
												v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v183 = m.T0[v182].(func(*base.Module, int32, int32) int32)(m, l0, v173)
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return int32(0)
												} else {
													v202 = v183
													m.G0 = v9 + int32(16)
													return v202
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
		v188 = m.ExcPending
		if v188 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_tuplestore_gettuple_3), int32(0))
			mBase = m.M
			v192 = m.ExcPending
			if v192 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_tuplestore_gettuple_1), int32(1111), int32(_a_F_tuplestore_gettuple_2))
				mBase = m.M
				v197 = m.ExcPending
				if v197 != 0 {
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
