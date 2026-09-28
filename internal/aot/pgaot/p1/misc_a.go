package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AcquireDeletionLock(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v43 int32
	_ = v43
	var v60 int32
	_ = v60
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 == int32(1259) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if l1&int32(2) != 0 {
			F_LockRelationOid(m, v7, int32(4))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		} else {
			F_LockRelationOid(m, v7, int32(8))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v18 = int32(1)
		if v4 <= int32(3591) {
			if v4 <= int32(2670) {
				switch v4 - int32(1213) {
				case 0, 1, 19, 20, 47, 48, 49:
					v89 = v18
				case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
					v89 = int32(0)
				default:
					if base.Ui32(int32(2)) <= base.Ui32(v4-int32(2396)) {
						v89 = int32(0)
					} else {
						v89 = v18
					}
				}
			} else {
				v30 = v4 - int32(2671)
				if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v30))|base.B2i32(int32(1)<<(uint(v30)%32)&int32(226492515) == int32(0)) != 0 {
					if base.B2i32(base.Ui32(v4-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v4-int32(2846)) < base.Ui32(int32(2))) != 0 {
						v89 = v18
					} else {
						v89 = int32(0)
					}
				} else {
					v89 = v18
				}
			}
		} else {
			if v4 <= int32(_a_F_AcquireDeletionLock_0) {
				v43 = v4 - int32(_a_F_AcquireDeletionLock_1)
				if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v43))|base.B2i32(int32(1)<<(uint(v43)%32)&int32(963) == int32(0)) != 0 {
					if base.Ui32(v4-int32(3592)) < base.Ui32(int32(2)) {
						v89 = v18
					} else {
						if base.Ui32(int32(2)) <= base.Ui32(v4-int32(4060)) {
							v89 = int32(0)
						} else {
							v89 = v18
						}
					}
				} else {
					v89 = v18
				}
			} else {
				switch v4 - int32(_a_F_AcquireDeletionLock_2) {
				case 0, 1, 2, 3, 4, 59, 60:
					v89 = v18
				case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
					v89 = int32(0)
				default:
					if base.Ui32(v4-int32(_a_F_AcquireDeletionLock_3)) < base.Ui32(int32(3)) {
						v89 = v18
					} else {
						v60 = v4 - int32(_a_F_AcquireDeletionLock_4)
						if base.Ui32(int32(15)) < base.Ui32(v60) {
							v89 = int32(0)
						} else {
							if int32(1)<<(uint(v60)%32)&int32(_a_F_AcquireDeletionLock_5) != 0 {
								v89 = v18
							} else {
								v89 = int32(0)
							}
						}
					}
				}
			}
		}
		v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v89 != 0 {
			F_LockSharedObject(m, v91, v90, int32(8))
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return
			} else {
				return
			}
		} else {
			F_LockDatabaseObject(m, v91, v90, int32(8))
			mBase = m.M
			v97 = m.ExcPending
			if v97 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_AlignedAllocRealloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v11 = l0 - int32(8)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v18 = v11 - base.I32_wrap_i64(int64(base.Ui64(v12)>>(uint(int64(34))%64)))&int32(1073741822)
	v19 = F_GetMemoryChunkSpace(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = F_GetMemoryChunkContext(m, v18)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v29 = base.I32_wrap_i64(int64(base.Ui64(v12)>>(uint(int64(5))%64))) & int32(1073741823)
			v30 = F_MemoryContextAllocAligned(m, v23, l1, v29, l2)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				if v30 == int32(0) {
					v34 = F_MemoryContextAllocationFailure(m, v23, l1, l2)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						return v34
					}
				} else {
					v39 = v19 - v29 - int32(8)
					if base.Ui32(l1) < base.Ui32(v39) {
						v41 = l1
					} else {
						v41 = v39
					}
					if v41 != 0 {
						base.MemoryCopy(m, v30, l0, v41)
					} else {
					}
					F_pfree(m, v18)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						return v30
					}
				}
			}
		}
	}
}
func F_AllocSetGetChunkContext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	v5 = l0 - int32(8)
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	if v6&int64(16) != int64(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(32))))
		return v13
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v5-base.I32_wrap_i64(int64(base.Ui64(v6)>>(uint(int64(34))%64)))&int32(1073741822))))
		return v21
	}
}
func F_AllocSetRealloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = l0 - int32(8)
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	if v18&int64(16) != int64(0) {
		v24 = l0 - int32(32)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
		if v25 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v170 = m.ExcPending
			if v170 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v17
				F_errmsg_internal(m, int32(_a_F_AllocSetRealloc_0), v14)
				mBase = m.M
				v174 = m.ExcPending
				if v174 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_AllocSetRealloc_1), int32(1267), int32(_a_F_AllocSetRealloc_2))
					mBase = m.M
					v179 = m.ExcPending
					if v179 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
			if v28 != int32(482) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v170 = m.ExcPending
				if v170 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v17
					F_errmsg_internal(m, int32(_a_F_AllocSetRealloc_0), v14)
					mBase = m.M
					v174 = m.ExcPending
					if v174 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_AllocSetRealloc_1), int32(1267), int32(_a_F_AllocSetRealloc_2))
						mBase = m.M
						v179 = m.ExcPending
						if v179 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(20))))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(16))))
				if v33 != v36 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v170 = m.ExcPending
					if v170 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v17
						F_errmsg_internal(m, int32(_a_F_AllocSetRealloc_0), v14)
						mBase = m.M
						v174 = m.ExcPending
						if v174 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_AllocSetRealloc_1), int32(1267), int32(_a_F_AllocSetRealloc_2))
							mBase = m.M
							v179 = m.ExcPending
							if v179 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v40 = int32(0)
					if (base.B2i32(l2&int32(1) == v40)|base.B2i32(l1 < v40))&base.B2i32(base.Ui32(int32(1073741824)) <= base.Ui32(l1)) != 0 {
						F_MemoryContextSizeFailure(m, l1)
						mBase = m.M
						v181 = m.ExcPending
						if v181 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v53 = (l1+int32(7))&int32(-8) + int32(32)
						v54 = F_emscripten_builtin_realloc(m, v24, v53)
						mBase = m.M
						if v54 == int32(0) {
							v57 = F_MemoryContextAllocationFailure(m, v25, l1, l2)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								v156 = v57
								m.G0 = v14 + int32(16)
								return v156
							}
						} else {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v61 + (v53 + v24 - v33)
							v66 = v54 + v53
							*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v66
							*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v66
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
							if v69 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v54
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v54
							}
							v73 = v54 + int32(32)
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
							if v74 == int32(0) {
								v156 = v73
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v54
								v156 = v73
							}
							m.G0 = v14 + int32(16)
							return v156
						}
					}
				}
			}
		}
	} else {
		v82 = int32(8) << (uint(base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(5))%64)))) % 32)
		if base.Ui32(l1) <= base.Ui32(v82) {
			v156 = l0
			m.G0 = v14 + int32(16)
			return v156
		} else {
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v17-base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(34))%64)))&int32(1073741822))))
			v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+104))
			if base.Ui32(v91) < base.Ui32(l1) {
				v93 = F_AllocSetAllocLarge(m, v90, l1, l2)
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					v127 = v93
					if v127 != 0 {
						v147 = v127
						if v82 != 0 {
							base.MemoryCopy(m, v147, l0, v82)
						} else {
						}
						F_AllocSetFree(m, l0)
						mBase = m.M
						v153 = m.ExcPending
						if v153 != 0 {
							return int32(0)
						} else {
							v156 = v147
							m.G0 = v14 + int32(16)
							return v156
						}
					} else {
						v128 = F_MemoryContextAllocationFailure(m, v90, l1, l2)
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							v156 = v128
							m.G0 = v14 + int32(16)
							return v156
						}
					}
				}
			} else {
				if base.Ui32(int32(8)) < base.Ui32(l1) {
					v103 = int32(29) - base.I32_clz(l1-int32(1))
				} else {
					v103 = int32(0)
				}
				v106 = v90 + v103<<(uint(int32(2))%32)
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
				if v107 != 0 {
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v106)+48)) = v108
					v147 = v107 + int32(8)
					if v82 != 0 {
						base.MemoryCopy(m, v147, l0, v82)
					} else {
					}
					F_AllocSetFree(m, l0)
					mBase = m.M
					v153 = m.ExcPending
					if v153 != 0 {
						return int32(0)
					} else {
						v156 = v147
						m.G0 = v14 + int32(16)
						return v156
					}
				} else {
					v112 = int32(8)
					v113 = v112 << (uint(v103) % 32)
					v116 = *(*int32)(unsafe.Add(mBase, uint32(v90)+44))
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
					v118 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
					if base.Ui32(v113+v112) <= base.Ui32(v117-v118) {
						v131 = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v116)+12)) = v118 + v113 + v131
						*(*int64)(unsafe.Add(mBase, uint32(v118))) = base.I64_extend_i32_u(v103<<(uint(int32(5))%32)) | base.I64_extend_i32_u(v118-v116)<<(uint(int64(34))%64) | int64(3)
						v147 = v118 + v131
						if v82 != 0 {
							base.MemoryCopy(m, v147, l0, v82)
						} else {
						}
						F_AllocSetFree(m, l0)
						mBase = m.M
						v153 = m.ExcPending
						if v153 != 0 {
							return int32(0)
						} else {
							v156 = v147
							m.G0 = v14 + int32(16)
							return v156
						}
					} else {
						v121 = F_AllocSetAllocFromNewBlock(m, v90, l1, l2, v103)
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return int32(0)
						} else {
							v127 = v121
							if v127 != 0 {
								v147 = v127
								if v82 != 0 {
									base.MemoryCopy(m, v147, l0, v82)
								} else {
								}
								F_AllocSetFree(m, l0)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return int32(0)
								} else {
									v156 = v147
									m.G0 = v14 + int32(16)
									return v156
								}
							} else {
								v128 = F_MemoryContextAllocationFailure(m, v90, l1, l2)
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int32(0)
								} else {
									v156 = v128
									m.G0 = v14 + int32(16)
									return v156
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_AlterSubscription_refresh_seq(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
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
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v387 int32
	_ = v387
	var v388 int64
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(208)
	m.G0 = v16
	v20 = l0 + int32(24)
	v24 = v3
	v25 = v3
	v26 = v3
	v27 = v3
	v28 = v3
	v29 = int32(-1)
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v29 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v387 = int32(m.ExcTag)
	v388 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v387 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L7:
	;
	v36 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+188)) = v36
	if l1 == v36 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v148 = v24
	v149 = v25
	v150 = v26
	v151 = v27
	v152 = v28
	goto L9
L9:
	;
	if v152 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v24
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v24
	F_load_file(m, int32(_a_F_AlterSubscription_refresh_seq_0), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L16
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v24
	F_errmsg_internal(m, int32(_a_F_AlterSubscription_refresh_seq_1), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v24
	F_errfinish(m, int32(_a_F_AlterSubscription_refresh_seq_2), int32(1347), int32(_a_F_AlterSubscription_refresh_seq_3))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L3
L16:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	if v74 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v80 = v77 ^ int32(1)
	goto L19
L18:
	;
	v80 = int32(0)
	goto L19
L19:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh_seq[0]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v20
	v89 = int32(1)
	v95 = m.T0[v84].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l1, v89, v89, v80&v89, v81, v16+int32(188))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	if v95 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v20
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L6
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh_seq[1]))
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh_seq[2]))
	goto L28
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v20
	F_errcode(m, int32(100663808))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v20
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v16)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v114
	F_errmsg(m, int32(_a_F_AlterSubscription_refresh_seq_4), v16+int32(16))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v20
	F_errfinish(m, int32(_a_F_AlterSubscription_refresh_seq_2), int32(1360), int32(_a_F_AlterSubscription_refresh_seq_3))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	goto L3
L28:
	;
	v142 = v16 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v16 + int32(28)
	goto L31
L29:
	;
	v148 = v20
	v149 = v95
	v150 = v140
	v151 = v138
	v152 = int32(0)
	goto L9
L31:
	;
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh_seq[2])) = v16 + int32(32)
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh_seq[0]))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	v167 = m.T0[v162].(func(*base.Module, int32) int32)(m, v149)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh_seq[1])) = v151
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh_seq[2])) = v150
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh_seq[0]))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	m.T0[v361].(func(*base.Module, int32))(m, v149)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L6
	} else {
		goto L63
	}
L35:
	;
	if v167 <= int32(_a_F_AlterSubscription_refresh_seq_5) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L6
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	v211 = int32(0)
	F_check_publications_origin_sequences(m, v149, v205, int32(1), v204, v211, v211, v203)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L6
	} else {
		goto L43
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	F_errcode(m, int32(325))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	F_errmsg(m, int32(_a_F_AlterSubscription_refresh_seq_6), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	F_errfinish(m, int32(_a_F_AlterSubscription_refresh_seq_2), int32(1373), int32(_a_F_AlterSubscription_refresh_seq_3))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	goto L3
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh_seq[1])) = v151
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh_seq[2])) = v150
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh_seq[0]))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	m.T0[v221].(func(*base.Module, int32))(m, v149)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh_seq[1])) = v151
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh_seq[2])) = v150
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	F_logicalrep_worker_stop(m, int32(2), v232, int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	v246 = int32(0)
	v249 = F_GetSubscriptionRelations(m, v241, v246, int32(1), v246)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L6
	} else {
		goto L47
	}
L46:
	;
	m.G0 = v16 + int32(208)
	return
L47:
	;
	if v249 == int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v253 = int32(0)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	if v254 <= v253 {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	v264 = v253
	goto L50
L50:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v270+v264<<(uint(int32(2))%32))))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	F_UpdateSubscriptionRelState(m, v276, v275, int32(105), int64(0), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L6
	} else {
		goto L52
	}
L51:
	;
	goto L46
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	v292 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	if v292 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	v298 = F_get_rel_namespace(m, v275)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L6
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v336 = v264 + int32(1)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	if v336 < v337 {
		v264 = v336
		goto L50
	} else {
		goto L62
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	v304 = F_get_namespace_name(m, v298)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	v310 = F_get_rel_name(m, v275)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v304
	F_errmsg_internal(m, int32(_a_F_AlterSubscription_refresh_seq_7), v16)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	F_errfinish(m, int32(_a_F_AlterSubscription_refresh_seq_2), int32(1428), int32(_a_F_AlterSubscription_refresh_seq_3))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	goto L56
L62:
	;
	goto L51
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+196)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v16)+200)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+204)) = v148
	F_pg_re_throw(m)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	goto L5
L65:
	;
	v392 = int32(v388)
	m.G0 = v16
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v392)+4))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	if v16+int32(28) == v398 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	m.ExcPending = 1
	goto L74
L67:
	;
	if v402 != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	v402 = v400
	goto L70
L69:
	;
	v402 = int32(0)
	goto L70
L70:
	;
	goto L67
L71:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v16)+204))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v16)+200))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v16)+196))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v16)+192))
	v24 = v403
	v25 = v404
	v26 = v405
	v27 = v406
	v28 = v394
	v29 = v402
	goto L1
L72:
	;
	goto L73
L73:
	;
	F___wasm_longjmp(m, v395, v394)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	return
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_AppendSeconds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	v8 = l1 >> (uint(int32(31)) % 32)
	v10 = l1 ^ v8 - v8
	if l3 != 0 {
		v11 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v10)) == int32(0) {
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10<<(uint(int32(1))%32))+uint32(_c_F_AppendSeconds[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v22)
			v37 = l0 + int32(2)
		} else {
			v26 = F_pg_ultoa_n(m, v10, l0)
			mBase = m.M
			if v11 <= v26 {
				v37 = l0 + v26
			} else {
				v29 = l0 + v11
				if v26 != 0 {
					base.MemoryCopy(m, v29-v26, l0, v26)
				} else {
				}
				v32 = v11 - v26
				if v32 != 0 {
					base.MemoryFill(m, l0, int32(48), v32)
				} else {
				}
				v37 = v29
			}
		}
		v40 = v37
	} else {
		v38 = F_pg_ultoa_n(m, v10, l0)
		mBase = m.M
		v40 = v38 + l0
	}
	if l2 == int32(0) {
		return v40
	} else {
		v44 = int32(46)
		*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v44)
		v47 = l2 >> (uint(int32(31)) % 32)
		v49 = l2 ^ v47 - v47
		v51 = base.I32_div_s(v49, int32(10))
		v54 = v51*int32(-10) + v49
		if v54 != 0 {
			v56 = v54 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v40)+6)) = uint8(v56)
			v62 = v40 + int32(7)
		} else {
			v62 = v40 + int32(6)
		}
		v64 = base.I32_div_s(v49, int32(100))
		v67 = v64*int32(-10) + v51
		v68 = v54 | v67
		if v68 == int32(0) {
			v76 = v40 + int32(5)
		} else {
			v74 = v67 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v40)+5)) = uint8(v74)
			v76 = v62
		}
		v78 = base.I32_div_s(v49, int32(1000))
		v81 = v64 + v78*int32(-10)
		v82 = v68 | v81
		if v82 == int32(0) {
			v90 = v40 + int32(4)
		} else {
			v88 = v81 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v40)+4)) = uint8(v88)
			v90 = v76
		}
		v92 = base.I32_div_s(v49, int32(_a_F_AppendSeconds_0))
		v95 = v78 + v92*int32(-10)
		v96 = v82 | v95
		if v96 == int32(0) {
			v104 = v40 + int32(3)
		} else {
			v102 = v95 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v40)+3)) = uint8(v102)
			v104 = v90
		}
		v106 = base.I32_div_s(v49, int32(_a_F_AppendSeconds_1))
		v109 = v92 + v106*int32(-10)
		v110 = v96 | v109
		if v110 == int32(0) {
			v118 = v40 + int32(2)
		} else {
			v116 = v109 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)) = uint8(v116)
			v118 = v104
		}
		v120 = base.I32_div_s(v49, int32(_a_F_AppendSeconds_2))
		v123 = v120*int32(-10) + v106
		if v110|v123 == int32(0) {
			v132 = v40 + int32(1)
		} else {
			v130 = v123 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)) = uint8(v130)
			v132 = v118
		}
		if base.Ui32(int32(19)) <= base.Ui32(v106+int32(9)) {
			v138 = v40 + int32(1)
			v139 = F_pg_ultoa_n(m, v49, v138)
			mBase = m.M
			v141 = v139 + v138
		} else {
			v141 = v132
		}
		return v141
	}
}
func F_AtAbort_Portals(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = v6 + int32(12)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_AtAbort_Portals[0]))
	F_hash_seq_init(m, v9, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = F_hash_seq_search(m, v9)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v14
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v6 + int32(32)
	return
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+80))
	if v20 != int32(3) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v39 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AtAbort_Portals[1])))
	if v24&int32(1) == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = int32(5)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v31 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	m.T0[v31].(func(*base.Module, int32))(m, v19)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(0)
	goto L9
L14:
	;
	v70 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L30
	}
L15:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+85)))
	if v42 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v19)+80))
	if v43 == int32(2) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = int32(5)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v48 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	m.T0[v48].(func(*base.Module, int32))(m, v19)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v53 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(0)
	goto L22
L24:
	;
	F_ReleaseCachedPlan(m, v53, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v19)+80))
	if v61 == int32(3) {
		goto L14
	} else {
		goto L28
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = int64(0)
	goto L26
L28:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	F_MemoryContextDeleteChildren(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L14
L30:
	;
	if v70 != 0 {
		v16 = v70
		goto L7
	} else {
		goto L31
	}
L31:
	;
	goto L8
}
func F_AtProcExit_Buffers(m *base.Module, l0 int32, l1 int64) {
	var v4 int32
	_ = v4
	F_UnlockBuffers(m)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_AtProcExit_Twophase(m *base.Module, l0 int32, l1 int64) {
	var v4 int32
	_ = v4
	F_AtAbort_Twophase(m)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_AuxiliaryPidGetProc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryPidGetProc[0]))
	v15 = v2
	goto L6
L4:
	;
	return v33
L5:
	;
	v33 = v22 + int32(768)
	goto L4
L6:
	;
	v18 = v15 * int32(768)
	v19 = v11 + v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v20 == l0 {
		v33 = v19
		goto L4
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v22 = v11 + v18
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+780))
	if v23 == l0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v26 = v15 + int32(2)
	if v26 != int32(38) {
		v15 = v26
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
}
func F_a_cas(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v2 != 0 {
		v4 = v2
	} else {
		v4 = int32(1073741823)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v4
	return
}
func F_accumArrayResult(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	v3 = l2
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == int32(0) {
		v17 = F_AllocSetContextCreateInternal(m, l4, int32(_a_F_accumArrayResult_0), int32(0), int32(_a_F_accumArrayResult_1), int32(_a_F_accumArrayResult_2))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v22 = F_MemoryContextAlloc(m, v17, int32(32))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v22)+28)) = uint8(v24)
				*(*int32)(unsafe.Add(mBase, uint32(v22))) = v17
				*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = int32(64)
				v30 = F_MemoryContextAlloc(m, v17, int32(512))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v30
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
					v34 = F_MemoryContextAlloc(m, v17, v33)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v34
						F_get_typlenbyvalalign(m, l3, v22+int32(24), v22+int32(26), v22+int32(27))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = v22
							v50 = int32(_a_F_accumArrayResult_3)
							v51 = *(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0]))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
							*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v53
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							if v55 <= v56 {
								*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v55 << (uint(int32(1)) % 32)
								v62 = v55 << (uint(int32(4)) % 32)
								if base.Ui32(int32(1073741824)) <= base.Ui32(v62) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(261))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(1073741823)
											F_errmsg(m, int32(_a_F_accumArrayResult_4), v9)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_accumArrayResult_5), int32(_a_F_accumArrayResult_6), int32(_a_F_accumArrayResult_0))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
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
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
									v66 = F_repalloc(m, v65, v62)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v66
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
										v71 = F_repalloc(m, v69, v70)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v71
											if v3 != 0 {
												v86 = l1
												v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
												*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
												m.G0 = v9 + int32(16)
												return v48
											} else {
												v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
												if v75 != 0 {
													v86 = l1
													v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
													*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
													v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
													*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
													v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
													*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
													*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
													m.G0 = v9 + int32(16)
													return v48
												} else {
													v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+24)))
													if v76 == int32(-1) {
														v80 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(l1))
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															v86 = base.I64_extend_i32_u(v80)
															v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
															v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
															v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
															*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
															m.G0 = v9 + int32(16)
															return v48
														}
													} else {
														v84 = F_datumCopy(m, l1, int32(0), v76)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int32(0)
														} else {
															v86 = v84
															v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
															v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
															v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
															*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
															m.G0 = v9 + int32(16)
															return v48
														}
													}
												}
											}
										}
									}
								}
							} else {
								if v3 != 0 {
									v86 = l1
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
									*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
									*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
									m.G0 = v9 + int32(16)
									return v48
								} else {
									v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
									if v75 != 0 {
										v86 = l1
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
										*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
										m.G0 = v9 + int32(16)
										return v48
									} else {
										v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+24)))
										if v76 == int32(-1) {
											v80 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(l1))
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												v86 = base.I64_extend_i32_u(v80)
												v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
												*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
												m.G0 = v9 + int32(16)
												return v48
											}
										} else {
											v84 = F_datumCopy(m, l1, int32(0), v76)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int32(0)
											} else {
												v86 = v84
												v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
												*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
												m.G0 = v9 + int32(16)
												return v48
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
		v48 = l0
		v50 = int32(_a_F_accumArrayResult_3)
		v51 = *(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0]))
		v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
		*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v53
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
		if v55 <= v56 {
			*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v55 << (uint(int32(1)) % 32)
			v62 = v55 << (uint(int32(4)) % 32)
			if base.Ui32(int32(1073741824)) <= base.Ui32(v62) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(1073741823)
						F_errmsg(m, int32(_a_F_accumArrayResult_4), v9)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_accumArrayResult_5), int32(_a_F_accumArrayResult_6), int32(_a_F_accumArrayResult_0))
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
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
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
				v66 = F_repalloc(m, v65, v62)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v66
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
					v71 = F_repalloc(m, v69, v70)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v71
						if v3 != 0 {
							v86 = l1
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
							m.G0 = v9 + int32(16)
							return v48
						} else {
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
							if v75 != 0 {
								v86 = l1
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
								*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
								m.G0 = v9 + int32(16)
								return v48
							} else {
								v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+24)))
								if v76 == int32(-1) {
									v80 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(l1))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										v86 = base.I64_extend_i32_u(v80)
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
										*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
										m.G0 = v9 + int32(16)
										return v48
									}
								} else {
									v84 = F_datumCopy(m, l1, int32(0), v76)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										v86 = v84
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
										*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
										m.G0 = v9 + int32(16)
										return v48
									}
								}
							}
						}
					}
				}
			}
		} else {
			if v3 != 0 {
				v86 = l1
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
				v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
				*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
				m.G0 = v9 + int32(16)
				return v48
			} else {
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
				if v75 != 0 {
					v86 = l1
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
					*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
					m.G0 = v9 + int32(16)
					return v48
				} else {
					v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+24)))
					if v76 == int32(-1) {
						v80 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(l1))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v86 = base.I64_extend_i32_u(v80)
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
							m.G0 = v9 + int32(16)
							return v48
						}
					} else {
						v84 = F_datumCopy(m, l1, int32(0), v76)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							v86 = v84
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(3))%32)))) = v86
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*uint8)(unsafe.Add(mBase, uint32(v94+v95))) = uint8(v3)
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v98 + int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
							m.G0 = v9 + int32(16)
							return v48
						}
					}
				}
			}
		}
	}
}
func F_aclcopy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) <= v10 {
		v16 = v10<<(uint(int32(4))%32) + int32(24)
		v17 = F_palloc0(m, v16)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v17)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v16 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v10
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v31 == int32(0) {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v41 = (v34<<(uint(int32(3))%32) + int32(23)) & int32(-8)
			} else {
				v41 = v31
			}
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v44 = v42 << (uint(int32(4)) % 32)
			if v44 != 0 {
				base.MemoryCopy(m, v17+int32(24), l0+v41, v44)
			} else {
			}
			m.G0 = v8 + int32(16)
			return v17
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
			F_errmsg_internal(m, int32(_a_F_aclcopy_0), v8)
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_aclcopy_1), int32(446), int32(_a_F_aclcopy_2))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
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
func F_aclinsert(m *base.Module, l0 int32) int64 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			F_errmsg(m, int32(_a_F_aclinsert_0), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_aclinsert_1), int32(1620), int32(_a_F_aclinsert_2))
				v19 = m.ExcPending
				if v19 != 0 {
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
func F_aclmembers(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	if l0 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v9 != 0 {
			F_check_acl(m, l0)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v21 = F_palloc(m, v18<<(uint(int32(3))%32))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v23 == int32(0) {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v33 = (v26<<(uint(int32(3))%32) + int32(23)) & int32(-8)
					} else {
						v33 = v23
					}
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if int32(0) < v34 {
						v41 = int32(0)
						v42 = int32(0)
						for {
							v49 = l0 + v33 + v41<<(uint(int32(4))%32)
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
							if v50 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v21+v42<<(uint(int32(2))%32)))) = v50
								v57 = v42 + int32(1)
							} else {
								v57 = v42
							}
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
							if v58 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v21+v57<<(uint(int32(2))%32)))) = v58
								v65 = v57 + int32(1)
							} else {
								v65 = v57
							}
							v67 = v41 + int32(1)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							if v67 < v68 {
								v41 = v67
								v42 = v65
								continue
							} else {
								break
							}
							break
						}
						F_pg_qsort(m, v21, v65, int32(4), int32(506))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v21
							if base.Ui32(int32(2)) <= base.Ui32(v65) {
								v79 = int32(0)
								v81 = int32(1)
								for {
									v87 = int32(2)
									v89 = v21 + v81<<(uint(v87)%32)
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v21+v79<<(uint(v87)%32))))
									if base.B2i32(base.Ui32(v94) < base.Ui32(v93))-base.B2i32(base.Ui32(v93) < base.Ui32(v94)) == int32(0) {
										v108 = v79
									} else {
										v101 = v79 + int32(1)
										if v101 == v81 {
											v108 = v81
										} else {
											v106 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
											*(*int32)(unsafe.Add(mBase, uint32(v21+v101<<(uint(int32(2))%32)))) = v106
											v108 = v101
										}
									}
									v111 = v81 + int32(1)
									if v111 != v65 {
										v79 = v108
										v81 = v111
										continue
									} else {
										break
									}
									break
								}
								v123 = v108 + int32(1)
							} else {
								v123 = v65
							}
							return v123
						}
					} else {
						F_pg_qsort(m, v21, int32(0), int32(4), int32(506))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v21
							return int32(0)
						}
					}
				}
			}
		} else {
			v10 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v10
			return v10
		}
	} else {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v10
		return v10
	}
}
func F_acquire_sample_rows(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 float64
	_ = v106
	var v109 float64
	_ = v109
	var v113 float64
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 float64
	_ = v170
	var v171 float64
	_ = v171
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
	var v194 int32
	_ = v194
	var v200 float64
	_ = v200
	var v201 float64
	_ = v201
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v218 float64
	_ = v218
	var v230 float64
	_ = v230
	var v251 float64
	_ = v251
	var v255 float64
	_ = v255
	var v257 float64
	_ = v257
	var v264 float64
	_ = v264
	var v265 float64
	_ = v265
	var v268 float64
	_ = v268
	var v276 float64
	_ = v276
	var v277 float64
	_ = v277
	var v279 float64
	_ = v279
	var v282 float64
	_ = v282
	var v284 float64
	_ = v284
	var v285 float64
	_ = v285
	var v286 float64
	_ = v286
	var v288 float64
	_ = v288
	var v289 float64
	_ = v289
	var v291 int32
	_ = v291
	var v292 float64
	_ = v292
	var v296 float64
	_ = v296
	var v308 float64
	_ = v308
	var v313 float64
	_ = v313
	var v314 float64
	_ = v314
	var v315 float64
	_ = v315
	var v319 float64
	_ = v319
	var v321 float64
	_ = v321
	var v323 float64
	_ = v323
	var v326 float64
	_ = v326
	var v329 float64
	_ = v329
	var v335 float64
	_ = v335
	var v336 int32
	_ = v336
	var v337 float64
	_ = v337
	var v340 float64
	_ = v340
	var v344 float64
	_ = v344
	var v345 float64
	_ = v345
	var v349 float64
	_ = v349
	var v357 float64
	_ = v357
	var v358 float64
	_ = v358
	var v361 float64
	_ = v361
	var v371 float64
	_ = v371
	var v393 float64
	_ = v393
	var v396 float64
	_ = v396
	var v398 float64
	_ = v398
	var v399 float64
	_ = v399
	var v402 float64
	_ = v402
	var v410 float64
	_ = v410
	var v430 float64
	_ = v430
	var v438 float64
	_ = v438
	var v444 float64
	_ = v444
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v465 float64
	_ = v465
	var v467 float64
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v486 int32
	_ = v486
	var v492 float64
	_ = v492
	var v493 float64
	_ = v493
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v556 int32
	_ = v556
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 float64
	_ = v583
	var v585 float64
	_ = v585
	var v586 float64
	_ = v586
	var v588 float64
	_ = v588
	var v590 float64
	_ = v590
	var v593 float64
	_ = v593
	var v599 float64
	_ = v599
	var v602 float64
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 float64
	_ = v609
	var v610 float64
	_ = v610
	var v612 float64
	_ = v612
	var v616 int32
	_ = v616
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	v7 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(128)
	m.G0 = v21
	v23 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+120)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v21)+112)) = v23
	v29 = v21 + int32(80)
	v31 = F_RelationGetNumberOfBlocksInFork(m, l0, v7)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v36 = Fn14349(m, int64(32))
	mBase = m.M
	goto L3
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v31
	F_pg_prng_seed(m, v21+int32(96), base.I64_extend_i32_u(v36))
	mBase = m.M
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if base.Ui32(v45) < base.Ui32(v46) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[0]))
	if v52 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v48 = v45
	goto L7
L6:
	;
	v48 = v46
	goto L7
L7:
	;
	goto L4
L8:
	;
	v98 = v21 + int32(64)
	v99 = F_pg_prng_uint32(m)
	mBase = m.M
	F_pg_prng_seed(m, v98, base.I64_extend_i32_u(v99))
	mBase = m.M
	goto L13
L9:
	;
	goto L8
L10:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_acquire_sample_rows[1])))
	if v56&int32(1) == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v61 = int32(_a_F_acquire_sample_rows_0)
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2]))
	v64 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2])) = v63 + v64
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v67 + v64
	v71 = int32(0)
	v73 = int32(_a_F_acquire_sample_rows_1)
	v74 = base.AtomicRmwOr32(m, v71, v73, v71)
	*(*int64)(unsafe.Add(mBase, uint32(v52+int32(8))+232)) = base.I64_extend_i32_u(v48)
	v82 = base.AtomicRmwOr32(m, v71, v73, v71)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v83 + v64
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2])) = v89 - v64
	goto L9
L12:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[3]))
	if v116 != 0 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v106 = F_pg_prng_double(m, v98)
	mBase = m.M
	if base.F64_eq(v106, float64(0)) != 0 {
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v109 = F_log(m, v106)
	mBase = m.M
	v113 = F_exp(m, base.F64_div(base.F64_neg(v109), base.F64_convert_i32_s(l3)))
	mBase = m.M
	*(*float64)(unsafe.Add(mBase, uint32(v21+int32(56)))) = v113
	goto L12
L15:
	;
	goto L14
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L114
	}
L17:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_acquire_sample_rows[4])))
	if v118&int32(1) == int32(0) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v123 = int32(0)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v130 = m.T0[v129].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v123, v123, v123, v123, int32(32))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v133 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[5]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v139 = int32(0)
	v144 = F_read_stream_begin_relation(m, int32(9), v137, v138, v139, int32(536), v21+int32(80), v139)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+188))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+132))
	v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, v130, v144)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v149 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v164 = v7
	v169 = v7
	v170 = float64(-1)
	v171 = float64(0)
	goto L28
L26:
	;
	v556 = v7
	goto L27
L27:
	;
	F_read_stream_end(m, v144)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L97
	}
L28:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	v556 = v486
	goto L27
L30:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+188))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+136))
	v183 = m.T0[v182].(func(*base.Module, int32, int32, int32, int32) int32)(m, v130, v21+int32(120), v21+int32(112), v133)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v183 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v194 = v164
	v200 = v170
	v201 = v171
	goto L35
L33:
	;
	v486 = v164
	v492 = v170
	v493 = v171
	goto L34
L34:
	;
	v497 = v169 + int32(1)
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[0]))
	if v501 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L35:
	;
	if v194 < l3 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v486 = v463
	v492 = v465
	v493 = v467
	goto L34
L37:
	;
	v467 = base.F64_add(v201, float64(1))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)+188))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)+136))
	v475 = m.T0[v474].(func(*base.Module, int32, int32, int32, int32) int32)(m, v130, v21+int32(120), v21+int32(112), v133)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L89
	}
L38:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+44))
	v209 = m.T0[v208].(func(*base.Module, int32) int32)(m, v133)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if base.F64_lt(v200, float64(0)) != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2+v194<<(uint(int32(2))%32)))) = v209
	v463 = v194 + int32(1)
	v465 = v200
	goto L37
L42:
	;
	v217 = v21 + int32(56)
	v218 = float64(0)
	v230 = base.F64_convert_i32_s(l3)
	if base.F64_ge(base.F64_mul(v230, float64(22)), v201) != 0 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v438 = v200
	goto L44
L44:
	;
	if base.F64_le(v438, float64(0)) != 0 {
		goto L80
	} else {
		goto L81
	}
L45:
	;
	v438 = v430
	goto L44
L46:
	;
	goto L45
L47:
	;
	goto L50
L48:
	;
	goto L49
L49:
	;
	v284 = float64(1)
	v285 = base.F64_add(v201, v284)
	v286 = base.F64_sub(v201, v230)
	v288 = base.F64_add(v286, v284)
	v289 = base.F64_div(v285, v288)
	v291 = v21 + int32(64)
	v292 = *(*float64)(unsafe.Add(mBase, uint32(v217)))
	v296 = v292
	goto L57
L50:
	;
	v251 = F_pg_prng_double(m, v21+int32(64))
	mBase = m.M
	if base.F64_eq(v251, float64(0)) != 0 {
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v255 = base.F64_add(v201, float64(1))
	v257 = base.F64_div(base.F64_sub(v255, v230), v255)
	if base.F64_gt(v257, v251) == int32(0) {
		v430 = v218
		goto L46
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v264 = v255
	v265 = v257
	v268 = v218
	goto L54
L54:
	;
	v276 = float64(1)
	v277 = base.F64_add(v268, v276)
	v279 = base.F64_add(v264, v276)
	v282 = base.F64_mul(v265, base.F64_div(base.F64_sub(v279, v230), v279))
	if base.F64_gt(v282, v251) != 0 {
		v264 = v279
		v265 = v282
		v268 = v277
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v430 = v277
	goto L46
L56:
	;
	goto L55
L57:
	;
	v308 = F_pg_prng_double(m, v291)
	mBase = m.M
	if base.F64_eq(v308, float64(0)) != 0 {
		goto L57
	} else {
		goto L59
	}
L58:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v217))) = v410
	v430 = v314
	goto L46
L59:
	;
	v313 = base.F64_mul(v201, base.F64_add(v296, float64(-1)))
	v314 = base.F64_floor(v313)
	v315 = base.F64_add(v288, v314)
	v319 = base.F64_add(v201, v313)
	v321 = F_log(m, base.F64_div(base.F64_mul(v315, base.F64_mul(v289, base.F64_mul(v289, v308))), v319))
	mBase = m.M
	v323 = F_exp(m, base.F64_div(v321, v230))
	mBase = m.M
	v326 = base.F64_div(base.F64_mul(v288, base.F64_div(v319, v315)), v201)
	if base.F64_le(v323, v326) != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L58
L61:
	;
	v410 = base.F64_div(v326, v323)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v329 = base.F64_add(v201, v314)
	v335 = base.F64_div(base.F64_mul(base.F64_add(v329, float64(1)), base.F64_div(base.F64_mul(v285, v308), v288)), v319)
	v336 = base.F64_lt(v230, v314)
	if v336 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v337 = v315
	goto L66
L65:
	;
	v337 = v285
	goto L66
L66:
	;
	if base.F64_le(v337, v329) != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if v336 != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v371 = v335
	goto L69
L69:
	;
	goto L76
L70:
	;
	v340 = v201
	goto L72
L71:
	;
	v340 = base.F64_add(v286, v314)
	goto L72
L72:
	;
	v344 = v329
	v345 = v340
	v349 = v335
	goto L73
L73:
	;
	v357 = base.F64_mul(v349, base.F64_div(v344, v345))
	v358 = float64(-1)
	v361 = base.F64_add(v344, v358)
	if base.F64_ge(v361, v337) != 0 {
		v344 = v361
		v345 = base.F64_add(v345, v358)
		v349 = v357
		goto L73
	} else {
		goto L75
	}
L74:
	;
	v371 = v357
	goto L69
L75:
	;
	goto L74
L76:
	;
	v393 = F_pg_prng_double(m, v291)
	mBase = m.M
	if base.F64_eq(v393, float64(0)) != 0 {
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v396 = F_log(m, v371)
	mBase = m.M
	v398 = F_exp(m, base.F64_div(v396, v230))
	mBase = m.M
	v399 = F_log(m, v393)
	mBase = m.M
	v402 = F_exp(m, base.F64_div(base.F64_neg(v399), v230))
	mBase = m.M
	if base.F64_le(v398, base.F64_div(v319, v201)) == int32(0) {
		v296 = v402
		goto L57
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	v410 = v402
	goto L60
L80:
	;
	goto L84
L81:
	;
	goto L82
L82:
	;
	v463 = v194
	v465 = base.F64_add(v438, float64(-1))
	goto L37
L83:
	;
	v451 = l2 + base.I32_trunc_sat_f64_s(base.F64_mul(v444, base.F64_convert_i32_s(l3)))<<(uint(int32(2))%32)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	F_pfree(m, v452)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L87
	}
L84:
	;
	v444 = F_pg_prng_double(m, v21-int32(-64))
	mBase = m.M
	if base.F64_eq(v444, float64(0)) != 0 {
		goto L84
	} else {
		goto L86
	}
L85:
	;
	goto L83
L86:
	;
	goto L85
L87:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+44))
	v457 = m.T0[v456].(func(*base.Module, int32) int32)(m, v133)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451))) = v457
	goto L82
L89:
	;
	if v475 != 0 {
		v194 = v463
		v200 = v465
		v201 = v467
		goto L35
	} else {
		goto L90
	}
L90:
	;
	goto L36
L91:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+188))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+132))
	v545 = m.T0[v544].(func(*base.Module, int32, int32) int32)(m, v130, v144)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L95
	}
L92:
	;
	goto L91
L93:
	;
	v505 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_acquire_sample_rows[1])))
	if v505&int32(1) == int32(0) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v510 = int32(_a_F_acquire_sample_rows_0)
	v512 = *(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2]))
	v513 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2])) = v512 + v513
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v501)))
	*(*int32)(unsafe.Add(mBase, uint32(v501))) = v516 + v513
	v520 = int32(0)
	v522 = int32(_a_F_acquire_sample_rows_1)
	v523 = base.AtomicRmwOr32(m, v520, v522, v520)
	*(*int64)(unsafe.Add(mBase, uint32(v501+int32(16))+232)) = base.I64_extend_i32_u(v497)
	v531 = base.AtomicRmwOr32(m, v520, v522, v520)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v501)))
	*(*int32)(unsafe.Add(mBase, uint32(v501))) = v532 + v513
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2])) = v538 - v513
	goto L92
L95:
	;
	if v545 != 0 {
		v164 = v486
		v169 = v497
		v170 = v492
		v171 = v493
		goto L28
	} else {
		goto L96
	}
L96:
	;
	goto L29
L97:
	;
	F_ExecDropSingleTupleTableSlot(m, v133)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)+188))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v570)+12))
	m.T0[v571].(func(*base.Module, int32))(m, v130)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	if l3 == v556 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	F_qsort_interruptible(m, l2, l3, int32(4), int32(537), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v21)+92))
	if v580 <= int32(0) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L102
L104:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v602
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = v599
	v606 = F_errstart(m, l1, int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L108
	}
L105:
	;
	v583 = float64(0)
	v599 = v583
	v602 = v583
	goto L104
L106:
	;
	goto L107
L107:
	;
	v585 = *(*float64)(unsafe.Add(mBase, uint32(v21)+112))
	v586 = base.F64_convert_i32_u(v580)
	v588 = base.F64_convert_i32_u(v31)
	v590 = float64(0.5)
	v593 = *(*float64)(unsafe.Add(mBase, uint32(v21)+120))
	v599 = base.F64_floor(base.F64_add(base.F64_mul(base.F64_div(v585, v586), v588), v590))
	v602 = base.F64_floor(base.F64_add(base.F64_mul(base.F64_div(v593, v586), v588), v590))
	goto L104
L108:
	;
	if v606 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v609 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	v610 = *(*float64)(unsafe.Add(mBase, uint32(v21)+120))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+16)) = v610
	v612 = *(*float64)(unsafe.Add(mBase, uint32(v21)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+24)) = v612
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v556
	*(*float64)(unsafe.Add(mBase, uint32(v21)+40)) = v609
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v21)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v616
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v608 + int32(4)
	F_errmsg(m, int32(_a_F_acquire_sample_rows_2), v21)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	m.G0 = v21 + int32(128)
	return v556
L112:
	;
	F_errfinish(m, int32(_a_F_acquire_sample_rows_3), int32(1416), int32(_a_F_acquire_sample_rows_4))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	F_errmsg_internal(m, int32(_a_F_acquire_sample_rows_5), int32(0))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_acquire_sample_rows_6), int32(931), int32(_a_F_acquire_sample_rows_7))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_addCompoundAffixFlagValue(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v8 = m.G0
	v10 = v8 - int32(1024)
	m.G0 = v10
	v13 = l1
	goto L1
L1:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if base.B2i32(base.Ui32(v19-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v19 == int32(32)) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v104 = F_pg_mblen_cstr(m, v13)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L13
	} else {
		goto L36
	}
L6:
	;
	v32 = v13
	v34 = v19
	v35 = v10
	goto L9
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L13
	} else {
		goto L32
	}
L9:
	;
	switch v34 {
	case 0, 9, 10, 11, 12, 13, 32:
		goto L11
	default:
		goto L12
	}
L10:
	;
	v47 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v35))) = uint8(v47)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v50 < v49 {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	goto L10
L12:
	;
	v38 = F_pg_mblen_cstr(m, v32)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if base.Ui32(v35) < base.Ui32(v10+int32(1024)-v38) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v38 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v44 = v35
	goto L17
L17:
	;
	v45 = v32 + v38
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v32 = v45
	v34 = v46
	v35 = v44
	goto L9
L18:
	;
	base.MemoryCopy(m, v35, v32, v38)
	goto L20
L19:
	;
	goto L20
L20:
	;
	v44 = v38 + v35
	goto L17
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v72 = F_MemoryContextStrdup(m, v71, v10)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L13
	} else {
		goto L31
	}
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v69 = v52
	goto L21
L23:
	;
	goto L24
L24:
	;
	if v49 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v67
	v69 = v67
	goto L21
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v49 << (uint(int32(1)) % 32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v59 = F_repalloc(m, v56, v49*int32(24))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L13
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(10)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v65 = F_MemoryContextAlloc(m, v63, int32(120))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L13
	} else {
		goto L30
	}
L29:
	;
	v67 = v59
	goto L25
L30:
	;
	v67 = v65
	goto L25
L31:
	;
	v76 = v69 + v70*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v72
	v79 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)) = uint8(v79)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v81 + v79
	m.G0 = v10 + int32(1024)
	return
L32:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	F_errmsg(m, int32(_a_F_addCompoundAffixFlagValue_0), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_addCompoundAffixFlagValue_1), int32(1098), int32(_a_F_addCompoundAffixFlagValue_2))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	v13 = v104 + v13
	goto L1
}
func F_addFamilyMember(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v55
	F_errmsg(m, int32(_a_F_addFamilyMember_0), v12)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L14
	} else {
		goto L24
	}
L2:
	;
	v89 = F_lappend(m, v14, l1)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L14
	} else {
		goto L23
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v17 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v20 = int32(0)
	if v20 < v17 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = v17
	goto L7
L6:
	;
	v23 = v20
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v31 = v3
	goto L8
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v25+v31<<(uint(int32(2))%32))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v39 != v24 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v78 = v31 + int32(1)
	if v78 != v23 {
		v31 = v78
		goto L8
	} else {
		goto L22
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v41 != v42 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v44 != v45 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v57 = F_format_type_be(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v60 = F_format_type_be(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	if v47 == int32(1) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v55
	F_errmsg(m, int32(_a_F_addFamilyMember_1), v12+int32(16))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_addFamilyMember_2), int32(1460), int32(_a_F_addFamilyMember_3))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	goto L9
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v89
	m.G0 = v12 + int32(32)
	return
L24:
	;
	F_errfinish(m, int32(_a_F_addFamilyMember_2), int32(1453), int32(_a_F_addFamilyMember_3))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_addFkConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32) {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+119)))
	switch v26 - int32(112) {
	case 0, 2:
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
		v50 = F_ConstraintNameIsUsed(m, int32(0), v49, l2)
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return
		} else {
			if v50 != 0 {
				v52 = int32(0)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l4)+48))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+68))
				v57 = F_ChooseConstraintName(m, l2, v52, int32(_a_F_addFkConstraint_0), v55, v52)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					v59 = v57
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					if v60 == int32(0) {
						v63 = F_pstrdup(m, v59)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v63
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)+48))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+68))
							v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
							v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)))
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)))
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
							v74 = int32(0)
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l5)+56))
							v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+87)))
							v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+88)))
							v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+86)))
							if l7 != 0 {
								v89 = int32(0)
								v90 = int32(1)
							} else {
								v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+119)))
								v89 = base.B2i32(v85 != int32(112))
								v90 = int32(0)
							}
							v91 = F_CreateConstraintEntry(m, v59, v67, int32(102), v69, v70, v71, v72, l7, v73, l10, l8, l8, v74, l6, v75, l9, l11, l12, l13, l8, v76, v77, l15, l14, v78, v74, v74, v74, base.B2i32(l7 == v74), v90, v89, l17, l16)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
								if l7 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = l7
									*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(2606)
									v104 = v23 + int32(4)
									if l1 != 0 {
										F_recordDependencyOn(m, l0, v104, int32(80))
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(1259)
											v110 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
											*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v110
											v117 = int32(83)
											F_recordDependencyOn(m, l0, v104, v117)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return
											} else {
												F_CommandCounterIncrement(m)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													m.G0 = v23 + int32(16)
													return
												}
											}
										}
									} else {
										v117 = int32(105)
										F_recordDependencyOn(m, l0, v104, v117)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return
											} else {
												m.G0 = v23 + int32(16)
												return
											}
										}
									}
								} else {
									F_CommandCounterIncrement(m)
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return
									} else {
										m.G0 = v23 + int32(16)
										return
									}
								}
							}
						}
					} else {
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)+48))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+68))
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
						v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)))
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
						v74 = int32(0)
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l5)+56))
						v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+87)))
						v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+88)))
						v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+86)))
						if l7 != 0 {
							v89 = int32(0)
							v90 = int32(1)
						} else {
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+119)))
							v89 = base.B2i32(v85 != int32(112))
							v90 = int32(0)
						}
						v91 = F_CreateConstraintEntry(m, v59, v67, int32(102), v69, v70, v71, v72, l7, v73, l10, l8, l8, v74, l6, v75, l9, l11, l12, l13, l8, v76, v77, l15, l14, v78, v74, v74, v74, base.B2i32(l7 == v74), v90, v89, l17, l16)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
							if l7 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = l7
								*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(2606)
								v104 = v23 + int32(4)
								if l1 != 0 {
									F_recordDependencyOn(m, l0, v104, int32(80))
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(1259)
										v110 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
										*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v110
										v117 = int32(83)
										F_recordDependencyOn(m, l0, v104, v117)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return
											} else {
												m.G0 = v23 + int32(16)
												return
											}
										}
									}
								} else {
									v117 = int32(105)
									F_recordDependencyOn(m, l0, v104, v117)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return
									} else {
										F_CommandCounterIncrement(m)
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return
										} else {
											m.G0 = v23 + int32(16)
											return
										}
									}
								}
							} else {
								F_CommandCounterIncrement(m)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return
								} else {
									m.G0 = v23 + int32(16)
									return
								}
							}
						}
					}
				}
			} else {
				v59 = l2
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				if v60 == int32(0) {
					v63 = F_pstrdup(m, v59)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v63
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)+48))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+68))
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
						v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)))
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
						v74 = int32(0)
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l5)+56))
						v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+87)))
						v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+88)))
						v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+86)))
						if l7 != 0 {
							v89 = int32(0)
							v90 = int32(1)
						} else {
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+119)))
							v89 = base.B2i32(v85 != int32(112))
							v90 = int32(0)
						}
						v91 = F_CreateConstraintEntry(m, v59, v67, int32(102), v69, v70, v71, v72, l7, v73, l10, l8, l8, v74, l6, v75, l9, l11, l12, l13, l8, v76, v77, l15, l14, v78, v74, v74, v74, base.B2i32(l7 == v74), v90, v89, l17, l16)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
							if l7 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = l7
								*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(2606)
								v104 = v23 + int32(4)
								if l1 != 0 {
									F_recordDependencyOn(m, l0, v104, int32(80))
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(1259)
										v110 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
										*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v110
										v117 = int32(83)
										F_recordDependencyOn(m, l0, v104, v117)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return
											} else {
												m.G0 = v23 + int32(16)
												return
											}
										}
									}
								} else {
									v117 = int32(105)
									F_recordDependencyOn(m, l0, v104, v117)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return
									} else {
										F_CommandCounterIncrement(m)
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return
										} else {
											m.G0 = v23 + int32(16)
											return
										}
									}
								}
							} else {
								F_CommandCounterIncrement(m)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return
								} else {
									m.G0 = v23 + int32(16)
									return
								}
							}
						}
					}
				} else {
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)+48))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+68))
					v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
					v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
					v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)))
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)))
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
					v74 = int32(0)
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l5)+56))
					v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+87)))
					v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+88)))
					v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+86)))
					if l7 != 0 {
						v89 = int32(0)
						v90 = int32(1)
					} else {
						v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+119)))
						v89 = base.B2i32(v85 != int32(112))
						v90 = int32(0)
					}
					v91 = F_CreateConstraintEntry(m, v59, v67, int32(102), v69, v70, v71, v72, l7, v73, l10, l8, l8, v74, l6, v75, l9, l11, l12, l13, l8, v76, v77, l15, l14, v78, v74, v74, v74, base.B2i32(l7 == v74), v90, v89, l17, l16)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
						if l7 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = l7
							*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(2606)
							v104 = v23 + int32(4)
							if l1 != 0 {
								F_recordDependencyOn(m, l0, v104, int32(80))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(1259)
									v110 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
									*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v110
									v117 = int32(83)
									F_recordDependencyOn(m, l0, v104, v117)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return
									} else {
										F_CommandCounterIncrement(m)
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return
										} else {
											m.G0 = v23 + int32(16)
											return
										}
									}
								}
							} else {
								v117 = int32(105)
								F_recordDependencyOn(m, l0, v104, v117)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return
								} else {
									F_CommandCounterIncrement(m)
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return
									} else {
										m.G0 = v23 + int32(16)
										return
									}
								}
							}
						} else {
							F_CommandCounterIncrement(m)
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return
							} else {
								m.G0 = v23 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = v36 + int32(4)
				F_errmsg(m, int32(_a_F_addFkConstraint_1), v23)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_addFkConstraint_2), int32(_a_F_addFkConstraint_3), int32(_a_F_addFkConstraint_4))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
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
func F_addFkRecurseReferencing(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v94 int32
	_ = v94
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v189 int32
	_ = v189
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v313 int32
	_ = v313
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v379 int32
	_ = v379
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v466 int32
	_ = v466
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v624 int32
	_ = v624
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	v20 = int32(0)
	v32 = m.G0
	v34 = v32 - int32(96)
	m.G0 = v34
	*(*int32)(unsafe.Add(mBase, uint32(v34)+92)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v34)+88)) = v20
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+119)))
	if v41 != int32(102) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
	if v44 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L10
	} else {
		goto L73
	}
L4:
	;
	m.G0 = v34 + int32(96)
	return
L5:
	;
	v223 = F_RelationGetPartitionDesc(m, l2, int32(1))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L10
	} else {
		goto L29
	}
L6:
	;
	if base.B2i32(l0 == int32(0))|l14 != 0 {
		goto L4
	} else {
		goto L12
	}
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	F_createForeignKeyCheckTriggers(m, v47, v48, l1, l5, l4, l16, l17, v34+int32(92), v34+int32(88))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v57 = v41
	goto L9
L9:
	;
	switch v57&int32(255) - int32(112) {
	case 0:
		goto L5
	default:
		goto L4
	case 2:
		goto L6
	}
L10:
	;
	return
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+119)))
	v57 = v56
	goto L9
L12:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)))
	if v65 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
	if v66 != int32(1) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v70 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v204 = F_palloc0(m, int32(32))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L10
	} else {
		goto L26
	}
L16:
	;
	v150 = F_palloc0(m, int32(144))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L10
	} else {
		goto L23
	}
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v73 <= int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	v94 = int32(0)
	goto L19
L19:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v76+v94<<(uint(int32(2))%32))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if v113 == v69 {
		v189 = v112
		goto L15
	} else {
		goto L21
	}
L20:
	;
	goto L16
L21:
	;
	v116 = v94 + int32(1)
	if v73 != v116 {
		v94 = v116
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v69
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+4)) = uint8(v156)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v159 = F_CreateTupleDescCopyConstr(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+8)) = v159
	*(*int64)(unsafe.Add(mBase, uint32(v150)+88)) = int64(0)
	v164 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+84)) = uint8(v164)
	v166 = int32(_a_F_addFkRecurseReferencing_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v150)+96)) = uint16(v166)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v169 = F_lappend(m, v168, v150)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v169
	v189 = v150
	goto L15
L26:
	;
	v206 = F_get_constraint_name(m, l5)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+4)) = int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v206
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v204)+20)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v204)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v204)+8)) = v211
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+84)))
	*(*int32)(unsafe.Add(mBase, uint32(v204)+24)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v204)+16)) = uint8(v215)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v189)+64))
	v219 = F_lappend(m, v218, v204)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+64)) = v219
	goto L4
L29:
	;
	v227 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	if int32(0) < v229 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v34)+88))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v34)+92))
	v264 = v20
	goto L34
L32:
	;
	goto L33
L33:
	;
	F_relation_close(m, v227, int32(3))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L10
	} else {
		goto L72
	}
L34:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v269+v264<<(uint(int32(2))%32))))
	v274 = F_table_open(m, v273, l15)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L10
	} else {
		goto L39
	}
L35:
	;
	goto L33
L36:
	;
	F_relation_close(m, v274, int32(0))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L10
	} else {
		goto L70
	}
L37:
	;
	v543 = int32(1)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v546 = v34 + int32(16)
	F_addFkConstraint(m, v34+int32(4), v543, v544, l1, v274, l3, l4, l5, l6, l7, v546, l9, l10, l11, l12, l13, v543, l18)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L10
	} else {
		goto L68
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L10
	} else {
		goto L64
	}
L39:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v274)+48))
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+118)))
	if v277 == int32(116) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+24)))
	if v280 == int32(0) {
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	F_CheckTableNotInUse(m, v274, int32(_a_F_addFkRecurseReferencing_1))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L10
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v274)+52))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v289 = F_build_attrmap_by_name(m, v286, v287, int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	if l6 <= int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v440 = F_RelationGetFKeyList(m, v274)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L10
	} else {
		goto L55
	}
L47:
	;
	v293 = int32(0)
	if l6 != int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v313 = v293
	v324 = v293
	goto L51
L49:
	;
	v379 = v293
	goto L50
L50:
	;
	v394 = int32(1)
	v395 = v379 << (uint(v394) % 32)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v401 = int32(*(*int16)(unsafe.Add(mBase, uint32(l8+v395))))
	v407 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v399+v401<<(uint(v394)%32)-int32(2)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v395+(v34+int32(16))))) = uint16(v407)
	goto L46
L51:
	;
	v328 = int32(1)
	v329 = v313 << (uint(v328) % 32)
	v331 = v34 + int32(16)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v335 = int32(*(*int16)(unsafe.Add(mBase, uint32(l8+v329))))
	v339 = int32(2)
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v333+v335<<(uint(v328)%32)-v339))))
	*(*uint16)(unsafe.Add(mBase, uint32(v329+v331))) = uint16(v341)
	v344 = v329 | v339
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v348 = int32(*(*int16)(unsafe.Add(mBase, uint32(l8+v344))))
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v346+v348<<(uint(v328)%32)-v339))))
	*(*uint16)(unsafe.Add(mBase, uint32(v331+v344))) = uint16(v354)
	v357 = v313 + v339
	v359 = v324 + v339
	if v359 != l6&int32(2147483646) {
		v313 = v357
		v324 = v359
		goto L51
	} else {
		goto L53
	}
L52:
	;
	if l6&int32(1) == int32(0) {
		goto L46
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v379 = v357
	goto L50
L55:
	;
	v442 = F_copyObjectImpl(m, v440)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	if v442 == int32(0) {
		goto L37
	} else {
		goto L57
	}
L57:
	;
	v446 = int32(0)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	if v447 <= v446 {
		goto L37
	} else {
		goto L58
	}
L58:
	;
	v466 = v446
	goto L59
L59:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v442)+12))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v481+v466<<(uint(int32(2))%32))))
	v488 = F_tryAttachPartitionForeignKey(m, l0, v485, v274, l5, l6, v34+int32(16), l7, l9, v237, v236, v227)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L10
	} else {
		goto L61
	}
L60:
	;
	goto L37
L61:
	;
	if v488 != 0 {
		goto L36
	} else {
		goto L62
	}
L62:
	;
	v491 = v466 + int32(1)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	if v491 < v492 {
		v466 = v491
		goto L59
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	F_errmsg(m, int32(_a_F_addFkRecurseReferencing_2), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_addFkRecurseReferencing_3), int32(_a_F_addFkRecurseReferencing_4), int32(_a_F_addFkRecurseReferencing_5))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	F_addFkRecurseReferencing(m, l0, l1, v274, l3, l4, v550, l6, l7, v546, l9, l10, l11, l12, l13, l14, l15, v237, v236, l18)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	goto L36
L70:
	;
	v588 = v264 + int32(1)
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	if v588 < v589 {
		v264 = v588
		goto L34
	} else {
		goto L71
	}
L71:
	;
	goto L35
L72:
	;
	goto L4
L73:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(_a_F_addFkRecurseReferencing_6), int32(0))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_addFkRecurseReferencing_3), int32(_a_F_addFkRecurseReferencing_7), int32(_a_F_addFkRecurseReferencing_8))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L10
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_addHLParsedLex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v244 int64
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v287 int32
	_ = v287
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var __phi325 int32
	_ = __phi325
	var v326 int32
	_ = v326
	var __phi326 int32
	_ = __phi326
	var v328 int32
	_ = v328
	var __phi328 int32
	_ = __phi328
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v369 int32
	_ = v369
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = l1 + int32(8)
	v19 = l2
	goto L4
L2:
	;
	goto L3
L3:
	;
	if l3 != 0 {
		goto L76
	} else {
		goto L77
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if int32(0) < v31 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v38 <= v37 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if l3 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v38 << (uint(int32(1)) % 32)
	v45 = F_repalloc(m, v34, v38<<(uint(int32(5))%32))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v49 = v34
	v50 = v37
	goto L11
L11:
	;
	v51 = int32(4)
	v53 = v50<<(uint(v51)%32) + v49
	v54 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+8)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v53))) = v54
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v58+v59<<(uint(v51)%32))+1)) = uint8(v31)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v64+v65<<(uint(v51)%32))+2)) = uint16(v35)
	v70 = F_palloc(m, v35)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L12
	} else {
		goto L14
	}
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v45
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v49 = v45
	v50 = v48
	goto L11
L14:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(4))%32))+8)) = v70
	if v35 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78+v79<<(uint(int32(4))%32))+8))
	base.MemoryCopy(m, v83, v36, v35)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v85 + int32(1)
	goto L8
L18:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	F_pfree(m, v19)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L12
	} else {
		goto L74
	}
L19:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v96 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v106 = l3
	v107 = v96
	v108 = v99
	goto L21
L21:
	;
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+2)))
	v117 = v108 + v114&int32(1)
	v118 = F_strlen(m, v107)
	mBase = m.M
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v120 <= v121+v122 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L18
L23:
	;
	v127 = v120
	v129 = v119
	goto L26
L24:
	;
	v156 = v119
	v157 = v121
	goto L25
L25:
	;
	v168 = v156 + v157<<(uint(int32(4))%32)
	v171 = int32(_a_F_addHLParsedLex_0)
	if v171 <= v117 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v127 << (uint(int32(1)) % 32)
	v144 = F_repalloc(m, v129, v127<<(uint(int32(5))%32))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L12
	} else {
		goto L28
	}
L27:
	;
	v156 = v144
	v157 = v148
	goto L25
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v144
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v147 <= v148+v149 {
		v127 = v147
		v129 = v144
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v174 = v171
	goto L32
L31:
	;
	v174 = v117
	goto L32
L32:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v168-int32(12)))) = uint16(v174)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v176 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v180 = v168 - int32(16)
	v182 = v168 - int32(4)
	v186 = v16
	v188 = int32(0)
	v189 = v176
	goto L36
L34:
	;
	goto L35
L35:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	if v287 != 0 {
		v106 = v106 + int32(8)
		v107 = v287
		v108 = v117
		goto L21
	} else {
		goto L73
	}
L36:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v198 != int32(1) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	v270 = v188 + int32(1)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v270 < v271 {
		v186 = v186 + int32(12)
		v188 = v270
		v189 = v271
		goto L36
	} else {
		goto L72
	}
L39:
	;
	v201 = int32(12)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	v209 = v204 & int32(4095)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+2)))
	if v209 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v235 != 0 {
		goto L38
	} else {
		goto L68
	}
L41:
	;
	if v210 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	if v118 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	v235 = int32(0)
	goto L40
L45:
	;
	goto L46
L46:
	;
	v215 = int32(0)
	if v215 < v118 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v218 = int32(-1)
	goto L49
L48:
	;
	v218 = v215
	goto L49
L49:
	;
	v235 = v218
	goto L40
L50:
	;
	v235 = base.B2i32(int32(0) < v209)
	goto L40
L51:
	;
	goto L52
L52:
	;
	if base.Ui32(v209) < base.Ui32(v118) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v224 = v209
	goto L55
L54:
	;
	v224 = v118
	goto L55
L55:
	;
	v225 = F_memcmp(m, v16+v189*v201+int32(base.Ui32(v204)>>(uint(v201)%32)), v107, v224)
	mBase = m.M
	if v210 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v235 = v233
	goto L40
L57:
	;
	if v225 != 0 {
		v233 = v225
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v225 != 0 {
		v233 = v225
		goto L56
	} else {
		goto L61
	}
L60:
	;
	v235 = base.B2i32(v118 < v209)
	goto L40
L61:
	;
	if v209 == v118 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v235 = int32(0)
	goto L40
L63:
	;
	goto L64
L64:
	;
	if v209 < v118 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v232 = int32(-1)
	goto L67
L66:
	;
	v232 = int32(1)
	goto L67
L67:
	;
	v233 = v232
	goto L56
L68:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	if v236 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v239 = int32(4)
	v241 = v237 + v238<<(uint(v239)%32)
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v180)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v242
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v180)))
	*(*int64)(unsafe.Add(mBase, uint32(v241))) = v244
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v246+v247<<(uint(v239)%32))+12)) = v186
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v256 = v252 + v253<<(uint(v239)%32)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v257 | int32(8)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v261 + int32(1)
	goto L38
L70:
	;
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v186
	goto L38
L72:
	;
	goto L37
L73:
	;
	goto L22
L74:
	;
	if v304 != 0 {
		v19 = v304
		goto L4
	} else {
		goto L75
	}
L75:
	;
	goto L5
L76:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v321 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	return
L79:
	;
	__phi325 = l3 + int32(4)
	__phi326 = l3
	__phi328 = v321
	v325 = __phi325
	v326 = __phi326
	v328 = __phi328
	goto L82
L80:
	;
	goto L81
L81:
	;
	F_pfree(m, l3)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L12
	} else {
		goto L89
	}
L82:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+2)))
	if v338&int32(1) != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	goto L81
L84:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v341 + int32(1)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v346 = v345
	goto L86
L85:
	;
	v346 = v328
	goto L86
L86:
	;
	F_pfree(m, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L12
	} else {
		goto L87
	}
L87:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	if v351 != 0 {
		__phi325 = v326 + int32(12)
		__phi326 = v326 + int32(8)
		__phi328 = v351
		v325 = __phi325
		v326 = __phi326
		v328 = __phi328
		goto L82
	} else {
		goto L88
	}
L88:
	;
	goto L83
L89:
	;
	goto L78
}
func F_addHyperLogLog(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v8 = int32(32) - v7
	v10 = v5 + int32(base.Ui32(l1)>>(uint(v8)%32))
	v11 = l1 << (uint(v7) % 32)
	if v11 != 0 {
		v18 = int32(32) - (base.I32_clz(v11) ^ int32(31))
		v19 = int32(255)
		if base.Ui32(v8&v19) < base.Ui32(v18&v19) {
			v24 = v8 + int32(1)
		} else {
			v24 = v18
		}
		v28 = v24
	} else {
		v28 = v8 + int32(1)
	}
	v30 = v28 & int32(255)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if base.Ui32(v31) < base.Ui32(v30) {
		v33 = v30
	} else {
		v33 = v31
	}
	*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v33)
	return
}
func F_add_dummy_return(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if v10 != 0 {
		v14 = F_palloc0(m, int32(32))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(0)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
			v19 = int32(1)
			v20 = v18 + v19
			*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v20
			*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v20
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v23
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v23
			v29 = F_list_make1_impl(m, v19, v7+int32(8))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(l0)+516)) = v14
				v33 = v29
				if v33 != 0 {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36<<(uint(int32(2))%32)-int32(4))))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					if v43 == int32(11) {
						m.G0 = v7 + int32(16)
						return
					} else {
						v47 = F_palloc0(m, int32(20))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(11)
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
							v53 = v51 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v53
							*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v53
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+472))
							*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v58
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
							v62 = F_lappend(m, v61, v47)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v62
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				} else {
					v47 = F_palloc0(m, int32(20))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(11)
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
						v53 = v51 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v53
						*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v53
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+472))
						*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v58
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
						v62 = F_lappend(m, v61, v47)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v62
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
		if v11 != 0 {
			v14 = F_palloc0(m, int32(32))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(0)
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
				v19 = int32(1)
				v20 = v18 + v19
				*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v20
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v23
				v29 = F_list_make1_impl(m, v19, v7+int32(8))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(l0)+516)) = v14
					v33 = v29
					if v33 != 0 {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36<<(uint(int32(2))%32)-int32(4))))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
						if v43 == int32(11) {
							m.G0 = v7 + int32(16)
							return
						} else {
							v47 = F_palloc0(m, int32(20))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(11)
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
								v53 = v51 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v53
								*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v53
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+472))
								*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v58
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
								v62 = F_lappend(m, v61, v47)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
									*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v62
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					} else {
						v47 = F_palloc0(m, int32(20))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(11)
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
							v53 = v51 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v53
							*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v53
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+472))
							*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v58
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
							v62 = F_lappend(m, v61, v47)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v62
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v33 = v12
			if v33 != 0 {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36<<(uint(int32(2))%32)-int32(4))))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
				if v43 == int32(11) {
					m.G0 = v7 + int32(16)
					return
				} else {
					v47 = F_palloc0(m, int32(20))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(11)
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
						v53 = v51 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v53
						*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v53
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+472))
						*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v58
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
						v62 = F_lappend(m, v61, v47)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v62
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			} else {
				v47 = F_palloc0(m, int32(20))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(11)
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
					v53 = v51 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v53
					*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v53
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+472))
					*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v58
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
					v62 = F_lappend(m, v61, v47)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v62
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_add_outer_joins_to_relids(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	if l2 == int32(0) {
		v310 = l1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v310
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v11 == int32(0) {
		v310 = l1
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v14 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = F_bms_add_member(m, l1, v11)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v23 = int32(0)
	if v22 == v23 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	return int32(0)
L8:
	;
	return v17
L9:
	;
	if v76 == int32(0) {
		v310 = l1
		goto L1
	} else {
		goto L23
	}
L10:
	;
	v76 = int32(1)
	goto L9
L11:
	;
	goto L12
L12:
	;
	if l1 == int32(0) {
		v69 = v23
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v76 = v69
	goto L9
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v33 < v32 {
		v69 = v23
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v35 = int32(1)
	if v32 <= v35 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v38 = v35
	goto L18
L17:
	;
	v38 = v32
	goto L18
L18:
	;
	v39 = int32(8)
	v44 = int32(0)
	goto L19
L19:
	;
	v51 = v44 << (uint(int32(2)) % 32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v22+v39+v51)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1+v39+v51)))
	v58 = v53 & (v55 ^ int32(-1))
	v60 = base.B2i32(v58 == int32(0))
	if v58 != 0 {
		v69 = v60
		goto L13
	} else {
		goto L21
	}
L20:
	;
	v69 = v60
	goto L13
L21:
	;
	v62 = v44 + int32(1)
	if v62 != v38 {
		v44 = v62
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v80 = F_bms_add_member(m, l1, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	if v82 == int32(0) {
		v310 = v80
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v85 = F_bms_copy(m, v82)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v87 == int32(0) {
		v310 = v80
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v90 <= int32(0) {
		v310 = v80
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v94 = int32(0)
	v95 = v80
	v99 = v85
	goto L29
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v94<<(uint(int32(2))%32))))
	if v106 == l2 {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v310 = v302
	goto L1
L31:
	;
	v306 = v94 + int32(1)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v306 < v307 {
		v94 = v306
		v95 = v302
		v99 = v303
		goto L29
	} else {
		goto L90
	}
L32:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)+24))
	if v108 == int32(0) {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
	if v111 != int32(1) {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v114 = F_bms_is_member(m, v108, v99)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	if v114 == int32(0) {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v106)+24))
	v119 = F_bms_is_member(m, v118, v95)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	if v119 != 0 {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L38
	}
L38:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v122 = int32(0)
	if v121 == v122 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v175 == int32(0) {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L53
	}
L40:
	;
	v175 = int32(1)
	goto L39
L41:
	;
	goto L42
L42:
	;
	if v95 == int32(0) {
		v168 = v122
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v175 = v168
	goto L39
L44:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v132 < v131 {
		v168 = v122
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v134 = int32(1)
	if v131 <= v134 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v137 = v134
	goto L48
L47:
	;
	v137 = v131
	goto L48
L48:
	;
	v138 = int32(8)
	v143 = int32(0)
	goto L49
L49:
	;
	v150 = v143 << (uint(int32(2)) % 32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v121+v138+v150)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v95+v138+v150)))
	v157 = v152 & (v154 ^ int32(-1))
	v159 = base.B2i32(v157 == int32(0))
	if v157 != 0 {
		v168 = v159
		goto L43
	} else {
		goto L51
	}
L50:
	;
	v168 = v159
	goto L43
L51:
	;
	v161 = v143 + int32(1)
	if v161 != v137 {
		v143 = v161
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	v179 = int32(0)
	if v178 == v179 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v232 == int32(0) {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L68
	}
L55:
	;
	v232 = int32(1)
	goto L54
L56:
	;
	goto L57
L57:
	;
	if v95 == int32(0) {
		v225 = v179
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v232 = v225
	goto L54
L59:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v189 < v188 {
		v225 = v179
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v191 = int32(1)
	if v188 <= v191 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v194 = v191
	goto L63
L62:
	;
	v194 = v188
	goto L63
L63:
	;
	v195 = int32(8)
	v200 = int32(0)
	goto L64
L64:
	;
	v207 = v200 << (uint(int32(2)) % 32)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v178+v195+v207)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v95+v195+v207)))
	v214 = v209 & (v211 ^ int32(-1))
	v216 = base.B2i32(v214 == int32(0))
	if v214 != 0 {
		v225 = v216
		goto L58
	} else {
		goto L66
	}
L65:
	;
	v225 = v216
	goto L58
L66:
	;
	v218 = v200 + int32(1)
	if v218 != v194 {
		v200 = v218
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v106)+36))
	v236 = int32(0)
	if v235 == v236 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v289 == int32(0) {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L83
	}
L70:
	;
	v289 = int32(1)
	goto L69
L71:
	;
	goto L72
L72:
	;
	if v95 == int32(0) {
		v282 = v236
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v289 = v282
	goto L69
L74:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v246 < v245 {
		v282 = v236
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v248 = int32(1)
	if v245 <= v248 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v251 = v248
	goto L78
L77:
	;
	v251 = v245
	goto L78
L78:
	;
	v252 = int32(8)
	v257 = int32(0)
	goto L79
L79:
	;
	v264 = v257 << (uint(int32(2)) % 32)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v235+v252+v264)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v95+v252+v264)))
	v271 = v266 & (v268 ^ int32(-1))
	v273 = base.B2i32(v271 == int32(0))
	if v271 != 0 {
		v282 = v273
		goto L73
	} else {
		goto L81
	}
L80:
	;
	v282 = v273
	goto L73
L81:
	;
	v275 = v257 + int32(1)
	if v275 != v251 {
		v257 = v275
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v106)+24))
	v293 = F_bms_add_member(m, v95, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	if l3 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v296 = F_lappend(m, v295, v106)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v106)+28))
	v300 = F_bms_add_members(m, v99, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L7
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v296
	goto L87
L89:
	;
	v302 = v293
	v303 = v300
	goto L31
L90:
	;
	goto L30
}
func F_add_rtes_to_flat_rtable(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	v2 = l1
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = F_palloc0(m, int32(16))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	if v40 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(384)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v28 = v26
	goto L8
L7:
	;
	v28 = int32(0)
	goto L8
L8:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+12)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v28
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v33 = F_lappend(m, v32, v18)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+44)) = v33
	goto L3
L10:
	;
	m.G0 = v13 + int32(16)
	return
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if int32(0) < v43 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v49 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+52))
	if v87 == int32(0) {
		goto L10
	} else {
		goto L24
	}
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v49<<(uint(int32(2))%32))))
	if v2 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L14
L17:
	;
	v73 = v49 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v73 < v74 {
		v49 = v73
		goto L15
	} else {
		goto L23
	}
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+56))
	F_add_rte_to_flat_rtable(m, v15, v69, v61)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L22
	}
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	switch v64 {
	case 0:
		goto L18
	case 1:
		goto L20
	default:
		goto L17
	}
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	if v65 == int32(0) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	goto L17
L23:
	;
	goto L16
L24:
	;
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v91 <= v90 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v97 = v90
	v98 = int32(1)
	goto L26
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105+v97<<(uint(int32(2))%32))))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	if v110 != int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L10
L28:
	;
	v173 = int32(1)
	v176 = v97 + v173
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v176 < v177 {
		v97 = v176
		v98 = v98 + v173
		goto L26
	} else {
		goto L53
	}
L29:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+20)))
	if v113 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v114) <= base.Ui32(v98) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116+v98<<(uint(int32(2))%32))))
	if v120 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)+148))
	if v123 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v15
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v109)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v127
	v133 = F_query_tree_walker_impl(m, v127, int32(882), v13+int32(8), int32(16))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v2 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L28
L37:
	;
	v166 = v123
	goto L39
L38:
	;
	v137 = F_fetch_upper_rel(m, v123, int32(7), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L40
	}
L39:
	;
	F_add_rtes_to_flat_rtable(m, v166, int32(1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L52
	}
L40:
	;
	v139 = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137)+44))
	if v141 == v139 {
		v162 = v139
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v162 == int32(0) {
		goto L28
	} else {
		goto L51
	}
L42:
	;
	goto L41
L43:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v145 = v144
	goto L44
L44:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if base.Ui32(int32(2)) <= base.Ui32(v149-int32(303)) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v162 = int32(1)
	goto L42
L46:
	;
	if v149 != int32(293) {
		v162 = v139
		goto L42
	} else {
		goto L49
	}
L47:
	;
	v145 = v148 + int32(72)
	goto L44
L48:
	;
	goto L45
L49:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v148)+72))
	if v156 != 0 {
		v162 = v139
		goto L42
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v120)+148))
	v166 = v165
	goto L39
L52:
	;
	goto L28
L53:
	;
	goto L27
}
func F_add_unique_group_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 float64
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 float32
	_ = v24
	var v26 float32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v62 float64
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 float64
	_ = v69
	var v72 int32
	_ = v72
	var v79 float64
	_ = v79
	var v82 float64
	_ = v82
	var v83 int32
	_ = v83
	var v88 float64
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 float64
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	v5 = int32(0)
	v8 = float64(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = v11 + int32(15)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v5)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v91 = F_remove_nulling_relids(m, l2, v89, int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L34
	} else {
		goto L35
	}
L2:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)))
	if v58 != 0 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
	v23 = v21 + v22
	v24 = *(*float32)(unsafe.Add(mBase, uint32(v23)+8))
	v26 = *(*float32)(unsafe.Add(mBase, uint32(v23)+16))
	v53 = base.F64_promote_f32(v26)
	v54 = base.F64_promote_f32(v24)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v28 == int32(16) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v53 = float64(2)
	v54 = v8
	goto L2
L7:
	;
	goto L8
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v32 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v39 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+84))
	if v35 != int32(5) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v53 = float64(-1)
	v54 = v8
	goto L2
L12:
	;
	v53 = float64(0)
	v54 = v8
	goto L2
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v42 != int32(6) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+8)))
	switch v46 - int32(_a_F_add_unique_group_var_0) {
	case 0:
		goto L15
	default:
		goto L12
	case 5:
		v53 = float64(-1)
		v54 = v8
		goto L2
	}
L15:
	;
	v53 = float64(1)
	v54 = v8
	goto L2
L16:
	;
	v59 = base.F64_neg(base.F64_sub(float64(1), v54))
	goto L18
L17:
	;
	v59 = v53
	goto L18
L18:
	;
	if base.F64_gt(v59, float64(0)) != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v62 = F_clamp_row_est(m, v59)
	mBase = m.M
	v88 = v62
	goto L1
L20:
	;
	goto L21
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v63 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v66 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v66)
	v88 = float64(200)
	goto L1
L23:
	;
	goto L24
L24:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v63)+128))
	if base.F64_le(v69, float64(0)) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v72 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v72)
	v88 = float64(200)
	goto L1
L26:
	;
	goto L27
L27:
	;
	if base.F64_lt(v59, float64(0)) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v79 = F_clamp_row_est(m, base.F64_mul(v69, base.F64_neg(v59)))
	mBase = m.M
	v88 = v79
	goto L1
L29:
	;
	goto L30
L30:
	;
	if base.F64_lt(v69, float64(200)) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v82 = F_clamp_row_est(m, v69)
	mBase = m.M
	v88 = v82
	goto L1
L32:
	;
	goto L33
L33:
	;
	v83 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v83)
	v88 = float64(200)
	goto L1
L34:
	;
	return int32(0)
L35:
	;
	if l1 == int32(0) {
		v139 = v5
		goto L37
	} else {
		goto L38
	}
L36:
	;
	m.G0 = v11 + int32(16)
	return v158
L37:
	;
	v144 = F_palloc(m, int32(24))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L34
	} else {
		goto L55
	}
L38:
	;
	v99 = int32(0)
	v102 = l1
	goto L39
L39:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	if v106 <= v99 {
		v139 = v102
		goto L37
	} else {
		goto L41
	}
L40:
	;
	v139 = v132
	goto L37
L41:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v99<<(uint(int32(2))%32))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v114 = F_equal(m, v91, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L34
	} else {
		goto L42
	}
L42:
	;
	if v114 != 0 {
		v158 = v102
		goto L36
	} else {
		goto L43
	}
L43:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v116 == v117 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v132 != 0 {
		v99 = v131 + int32(1)
		v102 = v132
		goto L39
	} else {
		goto L54
	}
L45:
	;
	v131 = v99
	v132 = v102
	goto L44
L46:
	;
	goto L47
L47:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v121 = F_exprs_known_equal(m, l0, v91, v119, int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L34
	} else {
		goto L48
	}
L48:
	;
	if v121 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v131 = v99
	v132 = v102
	goto L44
L50:
	;
	goto L51
L51:
	;
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v112)+8))
	if base.F64_le(v125, v88) != 0 {
		v158 = v102
		goto L36
	} else {
		goto L52
	}
L52:
	;
	v129 = F_list_delete_nth_cell(m, v102, v99)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L34
	} else {
		goto L53
	}
L53:
	;
	v131 = v99 - int32(1)
	v132 = v129
	goto L44
L54:
	;
	goto L40
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v91
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	*(*float64)(unsafe.Add(mBase, uint32(v144)+8)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v147
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+16)) = uint8(v150)
	v152 = F_lappend(m, v139, v144)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L34
	} else {
		goto L56
	}
L56:
	;
	v158 = v152
	goto L36
}
func F_adjust_view_column_set(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v3 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L29
	} else {
		goto L61
	}
L2:
	;
	if int32(0) <= v68 {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	v68 = base.I32_ctz(v54) | v55<<(uint(int32(5))%32)
	goto L2
L4:
	;
	v68 = int32(-2)
	goto L2
L5:
	;
	v19 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v22 <= v19 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v25 = l0 + int32(8)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v32 = v29 & int32(-1)
	if v32 != 0 {
		v54 = v32
		v55 = v19
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v33 = int32(1)
	if v33 == v22 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v37 = v33
	goto L9
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v25+v37<<(uint(int32(2))%32))))
	if v44 != 0 {
		v54 = v44
		v55 = v37
		goto L3
	} else {
		goto L11
	}
L10:
	;
	goto L4
L11:
	;
	v46 = v37 + int32(1)
	if v46 != v22 {
		v37 = v46
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v74 = v3
	v76 = v68
	goto L16
L14:
	;
	v238 = v3
	goto L15
L15:
	;
	m.G0 = v10 + int32(16)
	return v238
L16:
	;
	v79 = v76 - int32(7)
	if v79&int32(_a_F_adjust_view_column_set_0) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v238 = v173
	goto L15
L18:
	;
	if l0 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L19:
	;
	if l1 == int32(0) {
		v173 = v74
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v118 = base.I32_extend16_s(v79)
	if l1 != 0 {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v86 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v87 <= v86 {
		v173 = v74
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v92 = v86
	v93 = v74
	goto L24
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+v92<<(uint(int32(2))%32))))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+26)))
	if v102 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v173 = v113
	goto L18
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v105)+8)))
	v109 = F_bms_add_member(m, v93, v106+int32(7))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v113 = v93
	goto L28
L28:
	;
	v115 = v92 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v115 < v116 {
		v92 = v115
		v93 = v113
		goto L24
	} else {
		goto L31
	}
L29:
	;
	return int32(0)
L30:
	;
	v113 = v109
	goto L28
L31:
	;
	goto L25
L32:
	;
	if v156 == int32(0) {
		goto L1
	} else {
		goto L45
	}
L33:
	;
	goto L32
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v122 <= int32(0) {
		v156 = int32(0)
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v156 = int32(0)
	goto L33
L37:
	;
	v125 = int32(0)
	if v125 < v122 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v128 = v122
	goto L40
L39:
	;
	v128 = v125
	goto L40
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v133 = int32(0)
	goto L41
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v129+v133<<(uint(int32(2))%32))))
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+8)))
	if v142 == v118&int32(_a_F_adjust_view_column_set_0) {
		v156 = v141
		goto L33
	} else {
		goto L43
	}
L42:
	;
	goto L36
L43:
	;
	v145 = v133 + int32(1)
	if v145 != v128 {
		v133 = v145
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+26)))
	if v160 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v162 != int32(6) {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v165 = int32(*(*int16)(unsafe.Add(mBase, uint32(v161)+8)))
	v168 = F_bms_add_member(m, v74, v165+int32(7))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L29
	} else {
		goto L48
	}
L48:
	;
	v173 = v168
	goto L18
L49:
	;
	if int32(0) <= v232 {
		v74 = v173
		v76 = v232
		goto L16
	} else {
		goto L60
	}
L50:
	;
	v232 = base.I32_ctz(v218) | v219<<(uint(int32(5))%32)
	goto L49
L51:
	;
	v232 = int32(-2)
	goto L49
L52:
	;
	v183 = v76 + int32(1)
	v185 = int32(base.Ui32(v183) >> (uint(int32(5)) % 32))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v186 <= v185 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v189 = l0 + int32(8)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189+v185<<(uint(int32(2))%32))))
	v196 = v193 & (int32(-1) << (uint(v183) % 32))
	if v196 != 0 {
		v218 = v196
		v219 = v185
		goto L50
	} else {
		goto L54
	}
L54:
	;
	v198 = v185 + int32(1)
	if v198 == v186 {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v201 = v198
	goto L56
L56:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v189+v201<<(uint(int32(2))%32))))
	if v208 != 0 {
		v218 = v208
		v219 = v201
		goto L50
	} else {
		goto L58
	}
L57:
	;
	goto L51
L58:
	;
	v210 = v201 + int32(1)
	if v210 != v186 {
		v201 = v210
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	goto L17
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v118
	F_errmsg_internal(m, int32(_a_F_adjust_view_column_set_1), v10)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L29
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_adjust_view_column_set_2), int32(3178), int32(_a_F_adjust_view_column_set_3))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L29
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_af_6_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v5 - int32(1) {
	case 0:
		v8 = int32(0)
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v9 <= v10 {
			v163 = v8
			return v163
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v9-int32(1)))))
			if v16 != int32(105) {
				v163 = v8
				return v163
			} else {
				v20 = v9 - int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v20 <= v35 {
					v143 = int32(-1)
					v150 = v143
				} else {
					v52 = int32(1)
					v53 = v20 - v52
					v55 = int32(*(*int8)(unsafe.Add(mBase, uint32(v36+v53))))
					v57 = v55 & int32(255)
					if base.B2i32(v53 == v35)|base.B2i32(int32(0) <= v55) != 0 {
						v115 = v57
						v119 = v52
					} else {
						v64 = v57 & int32(63)
						v66 = v20 - int32(2)
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v66))))
						v70 = v68 << (uint(int32(6)) % 32)
						if base.B2i32(v66 != v35)&base.B2i32(base.Ui32(v68) < base.Ui32(int32(192))) == int32(0) {
							v115 = v70&int32(1984) | v64
							v119 = int32(2)
						} else {
							v83 = v70&int32(4032) | v64
							v85 = v20 - int32(3)
							v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v85))))
							if base.B2i32(v85 != v35)&base.B2i32(base.Ui32(v87) < base.Ui32(int32(224))) == int32(0) {
								v115 = v87<<(uint(int32(12))%32)&int32(_a_F_af_6_2_0) | v83
								v119 = int32(3)
							} else {
								v105 = int32(4)
								v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v36-v105))))
								v115 = v87<<(uint(int32(12))%32)&int32(_a_F_af_6_2_1) | v107&int32(7)<<(uint(int32(18))%32) | v83
								v119 = v105
							}
						}
					}
					if int32(246) < v115 {
						v150 = v119
					} else {
						v121 = v115 - int32(97)
						if v121 < int32(0) {
							v150 = v119
						} else {
							v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v121)>>(uint(int32(3))%32)))+uint32(_c_F_af_6_2[0]))))
							if int32(base.Ui32(v127)>>(uint(v121&int32(7))%32))&int32(1) == int32(0) {
								v150 = v119
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 - v119
								v143 = int32(0)
								v150 = v143
							}
						}
					}
				}
				return base.B2i32(v150 == int32(0))
			}
		}
	case 1:
		v157 = F_find_among_b(m, l0, int32(_a_F_af_6_2_2), int32(7), int32(0))
		mBase = m.M
		v160 = m.ExcPending
		if v160 != 0 {
			return int32(0)
		} else {
			v163 = base.B2i32(v157 != int32(0))
			return v163
		}
	default:
		v163 = int32(-1)
		return v163
	}
}
func F_akeys(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_akeys(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_anyarray_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anyarray_in_0), int32(154), int32(_a_F_anyarray_in_1), int32(_a_F_anyarray_in_2), int32(_a_F_anyarray_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_anycompatiblearray_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anycompatiblearray_in_0), int32(174), int32(_a_F_anycompatiblearray_in_1), int32(_a_F_anycompatiblearray_in_2), int32(_a_F_anycompatiblearray_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_anycompatiblenonarray_out(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anycompatiblenonarray_out_0), int32(377), int32(_a_F_anycompatiblenonarray_out_1), int32(_a_F_anycompatiblenonarray_out_2), int32(_a_F_anycompatiblenonarray_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_anycompatiblerange_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anycompatiblerange_in_0), int32(220), int32(_a_F_anycompatiblerange_in_1), int32(_a_F_anycompatiblerange_in_2), int32(_a_F_anycompatiblerange_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_anymultirange_out(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_multirange_out(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_anynonarray_out(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anynonarray_out_0), int32(375), int32(_a_F_anynonarray_out_1), int32(_a_F_anynonarray_out_2), int32(_a_F_anynonarray_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_anytimestamp_typmod_check(m *base.Module, l0 int32, l1 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14228(m, l0, l1, int32(_a_F_anytimestamp_typmod_check_0), int32(122), int32(_a_F_anytimestamp_typmod_check_1), int32(_a_F_anytimestamp_typmod_check_2), int32(129), int32(_a_F_anytimestamp_typmod_check_3))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_append_total_cost_compare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if v8 != v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v8 < v10 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v17 = int32(1)
	v18 = *(*float64)(unsafe.Add(mBase, uint32(v7)+56))
	v19 = *(*float64)(unsafe.Add(mBase, uint32(v9)+56))
	if base.F64_lt(v18, v19) != 0 {
		v84 = v17
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v15 = int32(1)
	goto L6
L5:
	;
	v15 = int32(-1)
	goto L6
L6:
	;
	return v15
L7:
	;
	return v84
L8:
	;
	if base.F64_gt(v18, v19) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(-1)
L10:
	;
	goto L11
L11:
	;
	v24 = *(*float64)(unsafe.Add(mBase, uint32(v7)+48))
	v25 = *(*float64)(unsafe.Add(mBase, uint32(v9)+48))
	if base.F64_lt(v24, v25) != 0 {
		v84 = v17
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if base.F64_gt(v24, v25) != 0 {
		v84 = int32(-1)
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v30 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v84 = v83
	goto L7
L15:
	;
	if v32 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	if v32 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v40 = int32(-1)
	goto L20
L19:
	;
	v40 = int32(0)
	goto L20
L20:
	;
	v83 = v40
	goto L14
L21:
	;
	v83 = int32(1)
	goto L14
L22:
	;
	goto L23
L23:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v44 != v45 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v45 < v44 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v51 = int32(8)
	v58 = v44 - int32(1)
	goto L31
L27:
	;
	v50 = int32(1)
	goto L29
L28:
	;
	v50 = int32(-1)
	goto L29
L29:
	;
	v83 = v50
	goto L14
L30:
	;
	if base.Ui32(v67) < base.Ui32(v65) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v63 = v58 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v30+v51+v63)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+(v32+v51))))
	if v65 != v67 {
		goto L30
	} else {
		goto L33
	}
L32:
	;
	v83 = int32(0)
	goto L14
L33:
	;
	v70 = v58 - int32(1)
	if int32(0) <= v70 {
		v58 = v70
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v77 = int32(1)
	goto L37
L36:
	;
	v77 = int32(-1)
	goto L37
L37:
	;
	v83 = v77
	goto L14
}
func F_apply_scanjoin_target_to_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v452 int32
	_ = v452
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	v7 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	if v19 == v7 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l4 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	v31 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v33 == v31 {
		v54 = v31
		goto L11
	} else {
		goto L12
	}
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	if v22 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+260))
	if v25 <= int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+276))
	if v28 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L3
L8:
	;
	return
L9:
	;
	v61 = v7
	goto L1
L10:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L20
	}
L11:
	;
	goto L10
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v37 = v36
	goto L13
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if base.Ui32(int32(2)) <= base.Ui32(v41-int32(303)) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v54 = int32(1)
	goto L11
L15:
	;
	if v41 != int32(293) {
		v54 = v31
		goto L11
	} else {
		goto L18
	}
L16:
	;
	v37 = v40 + int32(72)
	goto L13
L17:
	;
	goto L14
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v40)+72))
	if v48 != 0 {
		v54 = v31
		goto L11
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	if v54 != 0 {
		v61 = v7
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v57 = int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v58 {
	case 0, 2:
		goto L22
	default:
		v61 = v57
		goto L1
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = int32(0)
	v61 = v57
	goto L1
L23:
	;
	F_generate_useful_gather_paths(m, l0, l1, int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v61 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v67 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v67)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v67
	goto L25
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v78 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v73 {
	case 0, 2:
		goto L29
	default:
		goto L27
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(0)
	goto L27
L30:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v128 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L31:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v81 <= int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v91 = int32(0)
	goto L33
L33:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v102 = v99 + v91<<(uint(int32(2))%32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if l5 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L30
L35:
	;
	v111 = v91 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v111 < v112 {
		v91 = v111
		goto L33
	} else {
		goto L40
	}
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v104)+8)) = v105
	goto L35
L37:
	;
	goto L38
L38:
	;
	v107 = F_create_projection_path(m, l0, l1, v103, v77)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v107
	goto L35
L40:
	;
	goto L34
L41:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+38)))
	if v181 != 0 {
		goto L52
	} else {
		goto L53
	}
L42:
	;
	v131 = int32(0)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v132 <= v131 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v141 = v131
	goto L44
L44:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v152 = v149 + v141<<(uint(int32(2))%32)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	if l5 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L41
L46:
	;
	v163 = v141 + int32(1)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v163 < v164 {
		v141 = v163
		goto L44
	} else {
		goto L51
	}
L47:
	;
	v156 = F_create_projection_path(m, l0, l1, v153, v77)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L8
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+8)) = v160
	goto L46
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v156
	goto L46
L51:
	;
	goto L45
L52:
	;
	F_adjust_paths_for_srfs(m, l0, l1, l2, l3)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L8
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v184+v185<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v191
	if v61 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+280))
	if v193 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	goto L58
L58:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v485 != int32(1) {
		goto L124
	} else {
		goto L125
	}
L59:
	;
	if int32(0) <= v250 {
		goto L70
	} else {
		goto L71
	}
L60:
	;
	v250 = base.I32_ctz(v236) | v237<<(uint(int32(5))%32)
	goto L59
L61:
	;
	v250 = int32(-2)
	goto L59
L62:
	;
	v201 = int32(0)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	if v204 <= v201 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v207 = v193 + int32(8)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	v214 = v211 & int32(-1)
	if v214 != 0 {
		v236 = v214
		v237 = v201
		goto L60
	} else {
		goto L64
	}
L64:
	;
	v215 = int32(1)
	if v215 == v204 {
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v219 = v215
	goto L66
L66:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v207+v219<<(uint(int32(2))%32))))
	if v226 != 0 {
		v236 = v226
		v237 = v219
		goto L60
	} else {
		goto L68
	}
L67:
	;
	goto L61
L68:
	;
	v228 = v219 + int32(1)
	if v228 != v204 {
		v219 = v228
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v261 = v250
	v266 = v7
	goto L73
L71:
	;
	v468 = v7
	goto L72
L72:
	;
	F_add_paths_to_append_rel(m, l0, l1, v468)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L8
	} else {
		goto L123
	}
L73:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l1)+276))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267+v261<<(uint(int32(2))%32))))
	v272 = int32(0)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v271)+44))
	if v274 == v272 {
		v295 = v272
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v468 = v395
	goto L72
L75:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l1)+280))
	if v396 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L76:
	;
	if v295 != 0 {
		v395 = v266
		goto L75
	} else {
		goto L86
	}
L77:
	;
	goto L76
L78:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v278 = v277
	goto L79
L79:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	if base.Ui32(int32(2)) <= base.Ui32(v282-int32(303)) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v295 = int32(1)
	goto L77
L81:
	;
	if v282 != int32(293) {
		v295 = v272
		goto L77
	} else {
		goto L84
	}
L82:
	;
	v278 = v281 + int32(72)
	goto L79
L83:
	;
	goto L80
L84:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v281)+72))
	if v289 != 0 {
		v295 = v272
		goto L77
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v271)+8))
	v299 = F_find_appinfos_by_relids(m, l0, v296, v17+int32(12))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	v301 = int32(0)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v301 < v303 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v313 = v301
	v316 = v301
	goto L91
L89:
	;
	v348 = v301
	goto L90
L90:
	;
	F_pfree(m, v299)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L8
	} else {
		goto L97
	}
L91:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v320+v313<<(uint(int32(2))%32))))
	v325 = F_copy_pathtarget(m, v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L8
	} else {
		goto L93
	}
L92:
	;
	v348 = v332
	goto L90
L93:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v329 = F_adjust_appendrel_attrs(m, l0, v327, v328, v299)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v325)+4)) = v329
	v332 = F_lappend(m, v316, v325)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L8
	} else {
		goto L95
	}
L95:
	;
	v335 = v313 + int32(1)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v335 < v336 {
		v313 = v335
		v316 = v332
		goto L91
	} else {
		goto L96
	}
L96:
	;
	goto L92
L97:
	;
	F_apply_scanjoin_target_to_paths(m, l0, v271, v348, l3, l4, l5)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L8
	} else {
		goto L98
	}
L98:
	;
	v356 = int32(0)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v271)+44))
	if v358 == v356 {
		v379 = v356
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v379 != 0 {
		v395 = v266
		goto L75
	} else {
		goto L109
	}
L100:
	;
	goto L99
L101:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v358)+12))
	v362 = v361
	goto L102
L102:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	if base.Ui32(int32(2)) <= base.Ui32(v366-int32(303)) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v379 = int32(1)
	goto L100
L104:
	;
	if v366 != int32(293) {
		v379 = v356
		goto L100
	} else {
		goto L107
	}
L105:
	;
	v362 = v365 + int32(72)
	goto L102
L106:
	;
	goto L103
L107:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v365)+72))
	if v373 != 0 {
		v379 = v356
		goto L100
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v380 = F_lappend(m, v266, v271)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L8
	} else {
		goto L110
	}
L110:
	;
	v395 = v380
	goto L75
L111:
	;
	if int32(0) <= v452 {
		v261 = v452
		v266 = v395
		goto L73
	} else {
		goto L122
	}
L112:
	;
	v452 = base.I32_ctz(v438) | v439<<(uint(int32(5))%32)
	goto L111
L113:
	;
	v452 = int32(-2)
	goto L111
L114:
	;
	v403 = v261 + int32(1)
	v405 = int32(base.Ui32(v403) >> (uint(int32(5)) % 32))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	if v406 <= v405 {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v409 = v396 + int32(8)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v409+v405<<(uint(int32(2))%32))))
	v416 = v413 & (int32(-1) << (uint(v403) % 32))
	if v416 != 0 {
		v438 = v416
		v439 = v405
		goto L112
	} else {
		goto L116
	}
L116:
	;
	v418 = v405 + int32(1)
	if v418 == v406 {
		goto L113
	} else {
		goto L117
	}
L117:
	;
	v421 = v418
	goto L118
L118:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v409+v421<<(uint(int32(2))%32))))
	if v428 != 0 {
		v438 = v428
		v439 = v421
		goto L112
	} else {
		goto L120
	}
L119:
	;
	goto L113
L120:
	;
	v430 = v421 + int32(1)
	if v430 != v406 {
		v421 = v430
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	goto L74
L123:
	;
	goto L58
L124:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L8
	} else {
		goto L131
	}
L125:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(1)<<(uint(v488)%32)&int32(44) != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v496 = base.B2i32(base.Ui32(v488) <= base.Ui32(int32(5)))
	goto L128
L127:
	;
	v496 = int32(0)
	goto L128
L128:
	;
	if v496 != 0 {
		goto L124
	} else {
		goto L129
	}
L129:
	;
	F_generate_useful_gather_paths(m, l0, l1, int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L8
	} else {
		goto L130
	}
L130:
	;
	goto L124
L131:
	;
	m.G0 = v17 + int32(16)
	return
}
func F_apw_detach_shmem(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_apw_detach_shmem[0]))
	v7 = F_LWLockAcquire(m, v5, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_apw_detach_shmem[1]))
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_apw_detach_shmem[0]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
		if v10 == v13 {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(-1)
		} else {
		}
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
		if v10 == v17 {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(-1)
		} else {
		}
		F_LWLockRelease(m, v12)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			return
		}
	}
}
func F_apw_init_state(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v4 = F_LWLockNewTrancheId(m, int32(_a_F_apw_init_state_0))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_LWLockInitialize(m, l0, v4)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-1)
			return
		}
	}
}
func F_arraycontsel(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v27 int32
	_ = v27
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 float64
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int64
	_ = v129
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 float64
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 float64
	_ = v160
	var v161 int32
	_ = v161
	var v164 float64
	_ = v164
	var v166 float32
	_ = v166
	var v170 int32
	_ = v170
	var v176 float64
	_ = v176
	var v177 int32
	_ = v177
	var v182 float64
	_ = v182
	var v186 int32
	_ = v186
	var v191 float64
	_ = v191
	var v197 float64
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 float64
	_ = v203
	var v211 float64
	_ = v211
	var v220 int64
	_ = v220
	v9 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v24 = F_get_restriction_variable(m, v15, v16, v17, v12+int32(8), v12+int32(4), v12+int32(3))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int64(0)
	} else {
		if v24 == int32(0) {
			if v14 == int32(2750) {
				v34 = int64(4576918229304087675)
			} else {
				v34 = int64(4572414629676717179)
			}
			v220 = v34
			m.G0 = v12 + int32(112)
			return v220
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
			if v36 != int32(7) {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
				if v39 != 0 {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
					m.T0[v40].(func(*base.Module, int32))(m, v39)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						if v14 == int32(2750) {
							v47 = int64(4576918229304087675)
						} else {
							v47 = int64(4572414629676717179)
						}
						v220 = v47
						m.G0 = v12 + int32(112)
						return v220
					}
				} else {
					if v14 == int32(2750) {
						v47 = int64(4576918229304087675)
					} else {
						v47 = int64(4572414629676717179)
					}
					v220 = v47
					m.G0 = v12 + int32(112)
					return v220
				}
			} else {
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)))
				if v48 == int32(1) {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
					if v51 == int32(0) {
						v220 = v9
						m.G0 = v12 + int32(112)
						return v220
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
						m.T0[v54].(func(*base.Module, int32))(m, v51)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							v220 = v9
							m.G0 = v12 + int32(112)
							return v220
						}
					}
				} else {
					v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+3)))
					if v57 != 0 {
						v65 = v14
					} else {
						if v14 == int32(2751) {
							v65 = int32(2752)
						} else {
							if v14 == int32(2752) {
								v64 = int32(2751)
							} else {
								v64 = v14
							}
							v65 = v64
						}
					}
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
					v67 = F_get_base_element_type(m, v66)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int64(0)
					} else {
						if v67 == int32(0) {
							if v65 == int32(2750) {
								v191 = float64(0.01)
							} else {
								v191 = float64(0.005)
							}
							v197 = v191
							v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
							if v199 != 0 {
								v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
								m.T0[v200].(func(*base.Module, int32))(m, v199)
								mBase = m.M
								v202 = m.ExcPending
								if v202 != 0 {
									return int64(0)
								} else {
									v203 = float64(0)
									if base.F64_lt(v197, v203) != 0 {
										v211 = v203
									} else {
										if base.F64_gt(v197, float64(1)) == int32(0) {
											v211 = v197
										} else {
											v211 = float64(1)
										}
									}
									v220 = base.I64_reinterpret_f64(v211)
									m.G0 = v12 + int32(112)
									return v220
								}
							} else {
								v203 = float64(0)
								if base.F64_lt(v197, v203) != 0 {
									v211 = v203
								} else {
									if base.F64_gt(v197, float64(1)) == int32(0) {
										v211 = v197
									} else {
										v211 = float64(1)
									}
								}
								v220 = base.I64_reinterpret_f64(v211)
								m.G0 = v12 + int32(112)
								return v220
							}
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
							v72 = F_get_base_element_type(m, v71)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int64(0)
							} else {
								if v72 != v67 {
									if v65 == int32(2750) {
										v191 = float64(0.01)
									} else {
										v191 = float64(0.005)
									}
									v197 = v191
									v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
									if v199 != 0 {
										v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
										m.T0[v200].(func(*base.Module, int32))(m, v199)
										mBase = m.M
										v202 = m.ExcPending
										if v202 != 0 {
											return int64(0)
										} else {
											v203 = float64(0)
											if base.F64_lt(v197, v203) != 0 {
												v211 = v203
											} else {
												if base.F64_gt(v197, float64(1)) == int32(0) {
													v211 = v197
												} else {
													v211 = float64(1)
												}
											}
											v220 = base.I64_reinterpret_f64(v211)
											m.G0 = v12 + int32(112)
											return v220
										}
									} else {
										v203 = float64(0)
										if base.F64_lt(v197, v203) != 0 {
											v211 = v203
										} else {
											if base.F64_gt(v197, float64(1)) == int32(0) {
												v211 = v197
											} else {
												v211 = float64(1)
											}
										}
										v220 = base.I64_reinterpret_f64(v211)
										m.G0 = v12 + int32(112)
										return v220
									}
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
									v76 = *(*int64)(unsafe.Add(mBase, uint32(v75)+24))
									v78 = F_lookup_type_cache(m, v67, int32(64))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int64(0)
									} else {
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+108))
										if v80 == int32(0) {
											if v65 == int32(2750) {
												v87 = float64(0.01)
											} else {
												v87 = float64(0.005)
											}
											v197 = v87
											v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
											if v199 != 0 {
												v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
												m.T0[v200].(func(*base.Module, int32))(m, v199)
												mBase = m.M
												v202 = m.ExcPending
												if v202 != 0 {
													return int64(0)
												} else {
													v203 = float64(0)
													if base.F64_lt(v197, v203) != 0 {
														v211 = v203
													} else {
														if base.F64_gt(v197, float64(1)) == int32(0) {
															v211 = v197
														} else {
															v211 = float64(1)
														}
													}
													v220 = base.I64_reinterpret_f64(v211)
													m.G0 = v12 + int32(112)
													return v220
												}
											} else {
												v203 = float64(0)
												if base.F64_lt(v197, v203) != 0 {
													v211 = v203
												} else {
													if base.F64_gt(v197, float64(1)) == int32(0) {
														v211 = v197
													} else {
														v211 = float64(1)
													}
												}
												v220 = base.I64_reinterpret_f64(v211)
												m.G0 = v12 + int32(112)
												return v220
											}
										} else {
											v89 = F_pg_detoast_datum(m, base.I32_wrap_i64(v76))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int64(0)
											} else {
												v91 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
												if v91 == int32(0) {
													v170 = int32(0)
													v176 = F_mcelem_array_selec(m, v89, v78, v170, v170, v170, v170, v170, v170, v65)
													mBase = m.M
													v177 = m.ExcPending
													if v177 != 0 {
														return int64(0)
													} else {
														v182 = v176
														if v76 == base.I64_extend_i32_u(v89) {
															v197 = v182
															v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
															if v199 != 0 {
																v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
																m.T0[v200].(func(*base.Module, int32))(m, v199)
																mBase = m.M
																v202 = m.ExcPending
																if v202 != 0 {
																	return int64(0)
																} else {
																	v203 = float64(0)
																	if base.F64_lt(v197, v203) != 0 {
																		v211 = v203
																	} else {
																		if base.F64_gt(v197, float64(1)) == int32(0) {
																			v211 = v197
																		} else {
																			v211 = float64(1)
																		}
																	}
																	v220 = base.I64_reinterpret_f64(v211)
																	m.G0 = v12 + int32(112)
																	return v220
																}
															} else {
																v203 = float64(0)
																if base.F64_lt(v197, v203) != 0 {
																	v211 = v203
																} else {
																	if base.F64_gt(v197, float64(1)) == int32(0) {
																		v211 = v197
																	} else {
																		v211 = float64(1)
																	}
																}
																v220 = base.I64_reinterpret_f64(v211)
																m.G0 = v12 + int32(112)
																return v220
															}
														} else {
															F_pfree(m, v89)
															mBase = m.M
															v186 = m.ExcPending
															if v186 != 0 {
																return int64(0)
															} else {
																v197 = v182
																v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
																if v199 != 0 {
																	v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
																	m.T0[v200].(func(*base.Module, int32))(m, v199)
																	mBase = m.M
																	v202 = m.ExcPending
																	if v202 != 0 {
																		return int64(0)
																	} else {
																		v203 = float64(0)
																		if base.F64_lt(v197, v203) != 0 {
																			v211 = v203
																		} else {
																			if base.F64_gt(v197, float64(1)) == int32(0) {
																				v211 = v197
																			} else {
																				v211 = float64(1)
																			}
																		}
																		v220 = base.I64_reinterpret_f64(v211)
																		m.G0 = v12 + int32(112)
																		return v220
																	}
																} else {
																	v203 = float64(0)
																	if base.F64_lt(v197, v203) != 0 {
																		v211 = v203
																	} else {
																		if base.F64_gt(v197, float64(1)) == int32(0) {
																			v211 = v197
																		} else {
																			v211 = float64(1)
																		}
																	}
																	v220 = base.I64_reinterpret_f64(v211)
																	m.G0 = v12 + int32(112)
																	return v220
																}
															}
														}
													}
												} else {
													v96 = *(*int32)(unsafe.Add(mBase, uint32(v78)+108))
													v97 = F_statistic_proc_security_check(m, v12+int32(8), v96)
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return int64(0)
													} else {
														if v97 == int32(0) {
															v170 = int32(0)
															v176 = F_mcelem_array_selec(m, v89, v78, v170, v170, v170, v170, v170, v170, v65)
															mBase = m.M
															v177 = m.ExcPending
															if v177 != 0 {
																return int64(0)
															} else {
																v182 = v176
																if v76 == base.I64_extend_i32_u(v89) {
																	v197 = v182
																	v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
																	if v199 != 0 {
																		v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
																		m.T0[v200].(func(*base.Module, int32))(m, v199)
																		mBase = m.M
																		v202 = m.ExcPending
																		if v202 != 0 {
																			return int64(0)
																		} else {
																			v203 = float64(0)
																			if base.F64_lt(v197, v203) != 0 {
																				v211 = v203
																			} else {
																				if base.F64_gt(v197, float64(1)) == int32(0) {
																					v211 = v197
																				} else {
																					v211 = float64(1)
																				}
																			}
																			v220 = base.I64_reinterpret_f64(v211)
																			m.G0 = v12 + int32(112)
																			return v220
																		}
																	} else {
																		v203 = float64(0)
																		if base.F64_lt(v197, v203) != 0 {
																			v211 = v203
																		} else {
																			if base.F64_gt(v197, float64(1)) == int32(0) {
																				v211 = v197
																			} else {
																				v211 = float64(1)
																			}
																		}
																		v220 = base.I64_reinterpret_f64(v211)
																		m.G0 = v12 + int32(112)
																		return v220
																	}
																} else {
																	F_pfree(m, v89)
																	mBase = m.M
																	v186 = m.ExcPending
																	if v186 != 0 {
																		return int64(0)
																	} else {
																		v197 = v182
																		v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
																		if v199 != 0 {
																			v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
																			m.T0[v200].(func(*base.Module, int32))(m, v199)
																			mBase = m.M
																			v202 = m.ExcPending
																			if v202 != 0 {
																				return int64(0)
																			} else {
																				v203 = float64(0)
																				if base.F64_lt(v197, v203) != 0 {
																					v211 = v203
																				} else {
																					if base.F64_gt(v197, float64(1)) == int32(0) {
																						v211 = v197
																					} else {
																						v211 = float64(1)
																					}
																				}
																				v220 = base.I64_reinterpret_f64(v211)
																				m.G0 = v12 + int32(112)
																				return v220
																			}
																		} else {
																			v203 = float64(0)
																			if base.F64_lt(v197, v203) != 0 {
																				v211 = v203
																			} else {
																				if base.F64_gt(v197, float64(1)) == int32(0) {
																					v211 = v197
																				} else {
																					v211 = float64(1)
																				}
																			}
																			v220 = base.I64_reinterpret_f64(v211)
																			m.G0 = v12 + int32(112)
																			return v220
																		}
																	}
																}
															}
														} else {
															v101 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
															v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
															v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+22)))
															v110 = F_get_attstatsslot(m, v12+int32(76), v101, int32(4), int32(0), int32(3))
															mBase = m.M
															v111 = m.ExcPending
															if v111 != 0 {
																return int64(0)
															} else {
																if v110 != 0 {
																	if v65 != int32(2752) {
																		v126 = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v126
																		v129 = int64(0)
																		*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = v129
																		*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v129
																		*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v129
																		*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v129
																		v138 = v126
																		v139 = v126
																		v140 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
																		v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
																		v142 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
																		v143 = *(*int32)(unsafe.Add(mBase, uint32(v12)+100))
																		v144 = F_mcelem_array_selec(m, v89, v78, v140, v141, v142, v143, v139, v138, v65)
																		mBase = m.M
																		v145 = m.ExcPending
																		if v145 != 0 {
																			return int64(0)
																		} else {
																			F_free_attstatsslot(m, v12+int32(40))
																			mBase = m.M
																			v149 = m.ExcPending
																			if v149 != 0 {
																				return int64(0)
																			} else {
																				F_free_attstatsslot(m, v12+int32(76))
																				mBase = m.M
																				v153 = m.ExcPending
																				if v153 != 0 {
																					return int64(0)
																				} else {
																					v164 = v144
																					v166 = *(*float32)(unsafe.Add(mBase, uint32(v102+v103)+8))
																					v182 = base.F64_mul(v164, base.F64_sub(float64(1), base.F64_promote_f32(v166)))
																					if v76 == base.I64_extend_i32_u(v89) {
																						v197 = v182
																						v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
																						if v199 != 0 {
																							v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
																							m.T0[v200].(func(*base.Module, int32))(m, v199)
																							mBase = m.M
																							v202 = m.ExcPending
																							if v202 != 0 {
																								return int64(0)
																							} else {
																								v203 = float64(0)
																								if base.F64_lt(v197, v203) != 0 {
																									v211 = v203
																								} else {
																									if base.F64_gt(v197, float64(1)) == int32(0) {
																										v211 = v197
																									} else {
																										v211 = float64(1)
																									}
																								}
																								v220 = base.I64_reinterpret_f64(v211)
																								m.G0 = v12 + int32(112)
																								return v220
																							}
																						} else {
																							v203 = float64(0)
																							if base.F64_lt(v197, v203) != 0 {
																								v211 = v203
																							} else {
																								if base.F64_gt(v197, float64(1)) == int32(0) {
																									v211 = v197
																								} else {
																									v211 = float64(1)
																								}
																							}
																							v220 = base.I64_reinterpret_f64(v211)
																							m.G0 = v12 + int32(112)
																							return v220
																						}
																					} else {
																						F_pfree(m, v89)
																						mBase = m.M
																						v186 = m.ExcPending
																						if v186 != 0 {
																							return int64(0)
																						} else {
																							v197 = v182
																							v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
																							if v199 != 0 {
																								v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
																								m.T0[v200].(func(*base.Module, int32))(m, v199)
																								mBase = m.M
																								v202 = m.ExcPending
																								if v202 != 0 {
																									return int64(0)
																								} else {
																									v203 = float64(0)
																									if base.F64_lt(v197, v203) != 0 {
																										v211 = v203
																									} else {
																										if base.F64_gt(v197, float64(1)) == int32(0) {
																											v211 = v197
																										} else {
																											v211 = float64(1)
																										}
																									}
																									v220 = base.I64_reinterpret_f64(v211)
																									m.G0 = v12 + int32(112)
																									return v220
																								}
																							} else {
																								v203 = float64(0)
																								if base.F64_lt(v197, v203) != 0 {
																									v211 = v203
																								} else {
																									if base.F64_gt(v197, float64(1)) == int32(0) {
																										v211 = v197
																									} else {
																										v211 = float64(1)
																									}
																								}
																								v220 = base.I64_reinterpret_f64(v211)
																								m.G0 = v12 + int32(112)
																								return v220
																							}
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v116 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
																		v120 = F_get_attstatsslot(m, v12+int32(40), v116, int32(5), int32(0), int32(2))
																		mBase = m.M
																		v121 = m.ExcPending
																		if v121 != 0 {
																			return int64(0)
																		} else {
																			if v120 == int32(0) {
																				v126 = int32(0)
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v126
																				v129 = int64(0)
																				*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = v129
																				*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v129
																				*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v129
																				*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v129
																				v138 = v126
																				v139 = v126
																			} else {
																				v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
																				v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
																				v138 = v124
																				v139 = v125
																			}
																			v140 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
																			v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
																			v142 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
																			v143 = *(*int32)(unsafe.Add(mBase, uint32(v12)+100))
																			v144 = F_mcelem_array_selec(m, v89, v78, v140, v141, v142, v143, v139, v138, v65)
																			mBase = m.M
																			v145 = m.ExcPending
																			if v145 != 0 {
																				return int64(0)
																			} else {
																				F_free_attstatsslot(m, v12+int32(40))
																				mBase = m.M
																				v149 = m.ExcPending
																				if v149 != 0 {
																					return int64(0)
																				} else {
																					F_free_attstatsslot(m, v12+int32(76))
																					mBase = m.M
																					v153 = m.ExcPending
																					if v153 != 0 {
																						return int64(0)
																					} else {
																						v164 = v144
																						v166 = *(*float32)(unsafe.Add(mBase, uint32(v102+v103)+8))
																						v182 = base.F64_mul(v164, base.F64_sub(float64(1), base.F64_promote_f32(v166)))
																						if v76 == base.I64_extend_i32_u(v89) {
																							v197 = v182
																							v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
																							if v199 != 0 {
																								v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
																								m.T0[v200].(func(*base.Module, int32))(m, v199)
																								mBase = m.M
																								v202 = m.ExcPending
																								if v202 != 0 {
																									return int64(0)
																								} else {
																									v203 = float64(0)
																									if base.F64_lt(v197, v203) != 0 {
																										v211 = v203
																									} else {
																										if base.F64_gt(v197, float64(1)) == int32(0) {
																											v211 = v197
																										} else {
																											v211 = float64(1)
																										}
																									}
																									v220 = base.I64_reinterpret_f64(v211)
																									m.G0 = v12 + int32(112)
																									return v220
																								}
																							} else {
																								v203 = float64(0)
																								if base.F64_lt(v197, v203) != 0 {
																									v211 = v203
																								} else {
																									if base.F64_gt(v197, float64(1)) == int32(0) {
																										v211 = v197
																									} else {
																										v211 = float64(1)
																									}
																								}
																								v220 = base.I64_reinterpret_f64(v211)
																								m.G0 = v12 + int32(112)
																								return v220
																							}
																						} else {
																							F_pfree(m, v89)
																							mBase = m.M
																							v186 = m.ExcPending
																							if v186 != 0 {
																								return int64(0)
																							} else {
																								v197 = v182
																								v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
																								if v199 != 0 {
																									v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
																									m.T0[v200].(func(*base.Module, int32))(m, v199)
																									mBase = m.M
																									v202 = m.ExcPending
																									if v202 != 0 {
																										return int64(0)
																									} else {
																										v203 = float64(0)
																										if base.F64_lt(v197, v203) != 0 {
																											v211 = v203
																										} else {
																											if base.F64_gt(v197, float64(1)) == int32(0) {
																												v211 = v197
																											} else {
																												v211 = float64(1)
																											}
																										}
																										v220 = base.I64_reinterpret_f64(v211)
																										m.G0 = v12 + int32(112)
																										return v220
																									}
																								} else {
																									v203 = float64(0)
																									if base.F64_lt(v197, v203) != 0 {
																										v211 = v203
																									} else {
																										if base.F64_gt(v197, float64(1)) == int32(0) {
																											v211 = v197
																										} else {
																											v211 = float64(1)
																										}
																									}
																									v220 = base.I64_reinterpret_f64(v211)
																									m.G0 = v12 + int32(112)
																									return v220
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	v154 = int32(0)
																	v160 = F_mcelem_array_selec(m, v89, v78, v154, v154, v154, v154, v154, v154, v65)
																	mBase = m.M
																	v161 = m.ExcPending
																	if v161 != 0 {
																		return int64(0)
																	} else {
																		v164 = v160
																		v166 = *(*float32)(unsafe.Add(mBase, uint32(v102+v103)+8))
																		v182 = base.F64_mul(v164, base.F64_sub(float64(1), base.F64_promote_f32(v166)))
																		if v76 == base.I64_extend_i32_u(v89) {
																			v197 = v182
																			v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
																			if v199 != 0 {
																				v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
																				m.T0[v200].(func(*base.Module, int32))(m, v199)
																				mBase = m.M
																				v202 = m.ExcPending
																				if v202 != 0 {
																					return int64(0)
																				} else {
																					v203 = float64(0)
																					if base.F64_lt(v197, v203) != 0 {
																						v211 = v203
																					} else {
																						if base.F64_gt(v197, float64(1)) == int32(0) {
																							v211 = v197
																						} else {
																							v211 = float64(1)
																						}
																					}
																					v220 = base.I64_reinterpret_f64(v211)
																					m.G0 = v12 + int32(112)
																					return v220
																				}
																			} else {
																				v203 = float64(0)
																				if base.F64_lt(v197, v203) != 0 {
																					v211 = v203
																				} else {
																					if base.F64_gt(v197, float64(1)) == int32(0) {
																						v211 = v197
																					} else {
																						v211 = float64(1)
																					}
																				}
																				v220 = base.I64_reinterpret_f64(v211)
																				m.G0 = v12 + int32(112)
																				return v220
																			}
																		} else {
																			F_pfree(m, v89)
																			mBase = m.M
																			v186 = m.ExcPending
																			if v186 != 0 {
																				return int64(0)
																			} else {
																				v197 = v182
																				v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
																				if v199 != 0 {
																					v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
																					m.T0[v200].(func(*base.Module, int32))(m, v199)
																					mBase = m.M
																					v202 = m.ExcPending
																					if v202 != 0 {
																						return int64(0)
																					} else {
																						v203 = float64(0)
																						if base.F64_lt(v197, v203) != 0 {
																							v211 = v203
																						} else {
																							if base.F64_gt(v197, float64(1)) == int32(0) {
																								v211 = v197
																							} else {
																								v211 = float64(1)
																							}
																						}
																						v220 = base.I64_reinterpret_f64(v211)
																						m.G0 = v12 + int32(112)
																						return v220
																					}
																				} else {
																					v203 = float64(0)
																					if base.F64_lt(v197, v203) != 0 {
																						v211 = v203
																					} else {
																						if base.F64_gt(v197, float64(1)) == int32(0) {
																							v211 = v197
																						} else {
																							v211 = float64(1)
																						}
																					}
																					v220 = base.I64_reinterpret_f64(v211)
																					m.G0 = v12 + int32(112)
																					return v220
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
							}
						}
					}
				}
			}
		}
	}
}
func F_asin(m *base.Module, l0 float64) float64 {
	var v7 int64
	_ = v7
	var v12 int32
	_ = v12
	var v36 float64
	_ = v36
	var v71 float64
	_ = v71
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v107 float64
	_ = v107
	var v112 float64
	_ = v112
	var v117 float64
	_ = v117
	var v121 float64
	_ = v121
	var v130 float64
	_ = v130
	var v139 float64
	_ = v139
	var v143 float64
	_ = v143
	var v144 float64
	_ = v144
	v7 = base.I64_reinterpret_f64(l0)
	v12 = base.I32_wrap_i64(int64(base.Ui64(v7)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072693248)) <= base.Ui32(v12) {
		if base.I32_wrap_i64(v7)|(v12-int32(1072693248)) == int32(0) {
			return base.F64_add(base.F64_mul(l0, float64(1.5707963267948966)), float64(7.52316384526264e-37))
		} else {
			return base.F64_div(float64(0), base.F64_sub(l0, l0))
		}
	} else {
		if base.Ui32(v12) <= base.Ui32(int32(1071644671)) {
			if base.Ui32(v12+int32(-1048576)) < base.Ui32(int32(1044381696)) {
				v144 = l0
				return v144
			} else {
				v36 = base.F64_mul(l0, l0)
				return base.F64_add(base.F64_mul(l0, base.F64_div(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), float64(1)))), l0)
			}
		} else {
			v71 = float64(1)
			v75 = base.F64_mul(base.F64_sub(v71, base.F64_abs(l0)), float64(0.5))
			v76 = base.F64_sqrt(v75)
			v107 = base.F64_div(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), v71))
			if base.Ui32(int32(1072640819)) <= base.Ui32(v12) {
				v112 = base.F64_add(base.F64_mul(v76, v107), v76)
				v139 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v112, v112), float64(-6.123233995736766e-17)))
			} else {
				v117 = float64(0.7853981633974483)
				v121 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v76) & int64(-4294967296))
				v130 = base.F64_div(base.F64_sub(v75, base.F64_mul(v121, v121)), base.F64_add(v76, v121))
				v139 = base.F64_add(base.F64_sub(base.F64_sub(v117, base.F64_add(v121, v121)), base.F64_sub(base.F64_mul(base.F64_add(v76, v76), v107), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v130, v130)))), v117)
			}
			if v7 < int64(0) {
				v143 = base.F64_neg(v139)
			} else {
				v143 = v139
			}
			v144 = v143
			return v144
		}
	}
}
func F_assign_checkpoint_completion_target(m *base.Module, l0 float64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	*(*float64)(unsafe.Add(mBase, _c_F_assign_checkpoint_completion_target[0])) = l0
	v6 = int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_assign_checkpoint_completion_target[1]))
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_assign_checkpoint_completion_target[2]))
	v12 = base.I32_div_s(v10, int32(_a_F_assign_checkpoint_completion_target_0))
	v13 = base.I32_div_s(v8, v12)
	v18 = base.I32_trunc_sat_f64_s(base.F64_div(base.F64_convert_i32_s(v13), base.F64_add(l0, float64(1))))
	if v18 <= v6 {
		v21 = v6
	} else {
		v21 = v18
	}
	*(*int32)(unsafe.Add(mBase, _c_F_assign_checkpoint_completion_target[3])) = v21
	return
}
func F_assign_collations_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v682 int32
	_ = v682
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(96)
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+60)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v19
	v25 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v25
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v28 - int32(1) {
	case 0:
		goto L12
	default:
		goto L6
	case 5, 6, 7, 33, 55, 56, 57:
		goto L11
	case 8:
		goto L10
	case 10:
		goto L9
	case 13:
		goto L7
	case 24:
		goto L19
	case 30:
		goto L20
	case 31:
		goto L8
	case 35:
		goto L18
	case 36:
		goto L17
	case 53, 59, 62, 63, 64, 65, 105:
		goto L14
	case 54:
		goto L16
	case 61:
		goto L15
	case 66:
		goto L13
	}
L3:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	F_merge_collation_state(m, v855, v854, v857, v863, v864, l1)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L21
	} else {
		goto L231
	}
L4:
	;
	F_exprSetCollation(m, l0, int32(0))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L21
	} else {
		goto L230
	}
L5:
	;
	v768 = F_exprType(m, l0)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L21
	} else {
		goto L197
	}
L6:
	;
	v754 = F_expression_tree_walker_impl(m, l0, int32(518), v15+int32(48))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L21
	} else {
		goto L195
	}
L7:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+76)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v19
	v731 = v15 + int32(72)
	v732 = F_assign_collations_walker(m, v724, v731)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L21
	} else {
		goto L191
	}
L8:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v671 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L9:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v658 = F_assign_collations_walker(m, v655, v15+int32(48))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L21
	} else {
		goto L180
	}
L10:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	switch v267 - int32(104) {
	case 0:
		goto L86
	default:
		goto L85
	case 6:
		goto L84
	case 7:
		goto L87
	}
L11:
	;
	v262 = F_exprCollation(m, l0)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L21
	} else {
		goto L82
	}
L12:
	;
	v257 = F_expression_tree_walker_impl(m, l0, int32(518), v15+int32(48))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L21
	} else {
		goto L81
	}
L13:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v242 == int32(0) {
		goto L1
	} else {
		goto L78
	}
L14:
	;
	v240 = F_expression_tree_walker_impl(m, l0, int32(518), v15+int32(48))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L21
	} else {
		goto L77
	}
L15:
	;
	v193 = F_expression_tree_walker_impl(m, l0, int32(518), v15+int32(48))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L21
	} else {
		goto L64
	}
L16:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v165 = F_get_typcollation(m, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L21
	} else {
		goto L52
	}
L17:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v96 = int32(0)
	v103 = v3
	goto L33
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v52 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L19:
	;
	v44 = F_expression_tree_walker_impl(m, l0, int32(518), v15+int32(48))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L21
	} else {
		goto L23
	}
L20:
	;
	v34 = F_expression_tree_walker_impl(m, l0, int32(518), v15+int32(48))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v854 = int32(3)
	v855 = v39
	v857 = v38
	goto L3
L23:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v46 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v854 = v3
	v855 = int32(0)
	v857 = v25
	goto L3
L25:
	;
	goto L26
L26:
	;
	v51 = F_exprLocation(m, l0)
	mBase = m.M
	v854 = int32(1)
	v855 = v46
	v857 = v51
	goto L3
L27:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v55 <= int32(0) {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v59 = int32(0)
	goto L29
L29:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v59<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+76)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v19
	v83 = F_assign_collations_walker(m, v75, v15+int32(72))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L21
	} else {
		goto L31
	}
L30:
	;
	goto L1
L31:
	;
	v86 = v59 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v86 < v87 {
		v59 = v86
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v104 = int32(0)
	if v90 == v104 {
		v114 = v104
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v89 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v108 <= v96 {
		v114 = int32(0)
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v114 = v110 + v96<<(uint(int32(2))%32)
	goto L35
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	goto L1
L39:
	;
	goto L40
L40:
	;
	v119 = int32(0)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if base.B2i32(v114 == v119)|base.B2i32(v121 <= v96) == v119 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v126+v96<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v134
	v143 = F_list_make2_impl(m, v15+int32(12), v15+int32(8))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L21
	} else {
		goto L46
	}
L42:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v126 != 0 {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v103
	goto L1
L45:
	;
	goto L44
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+76)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v129
	v152 = F_assign_collations_walker(m, v143, v15+int32(72))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L21
	} else {
		goto L47
	}
L47:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	if v158 != int32(2) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v161 = v156
	goto L50
L49:
	;
	v161 = int32(0)
	goto L50
L50:
	;
	v162 = F_lappend_oid(m, v103, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L21
	} else {
		goto L51
	}
L51:
	;
	v96 = v96 + int32(1)
	v103 = v162
	goto L33
L52:
	;
	v170 = F_expression_tree_walker_impl(m, l0, int32(518), v15+int32(48))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L21
	} else {
		goto L53
	}
L53:
	;
	if v165 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_exprSetCollation(m, l0, v165)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L21
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v165 != int32(100) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v854 = v165
	v855 = v165
	v857 = v25
	goto L3
L58:
	;
	v179 = F_exprLocation(m, l0)
	mBase = m.M
	F_exprSetCollation(m, l0, v165)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L21
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v182 = int32(2)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	if v185 == v182 {
		goto L4
	} else {
		goto L62
	}
L61:
	;
	v854 = int32(1)
	v855 = v165
	v857 = v179
	goto L3
L62:
	;
	F_exprSetCollation(m, l0, v184)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L21
	} else {
		goto L63
	}
L63:
	;
	v854 = v185
	v855 = v184
	v857 = v183
	goto L3
L64:
	;
	v195 = int32(2)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	if v198 != v195 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v854 = v198
	v855 = v197
	v857 = v196
	goto L3
L66:
	;
	goto L67
L67:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v201 == int32(0) {
		v854 = v195
		v855 = v197
		v857 = v196
		goto L3
	} else {
		goto L68
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L21
	} else {
		goto L69
	}
L69:
	;
	F_errcode(m, int32(17432708))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L21
	} else {
		goto L70
	}
L70:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v212 = F_get_collation_name(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L21
	} else {
		goto L71
	}
L71:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v215 = F_get_collation_name(m, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L21
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v212
	F_errmsg(m, int32(_a_F_assign_collations_walker_0), v15+int32(16))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L21
	} else {
		goto L73
	}
L73:
	;
	F_errhint(m, int32(_a_F_assign_collations_walker_1), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L21
	} else {
		goto L74
	}
L74:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	F_parser_errposition(m, v228, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L21
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_assign_collations_walker_2), int32(480), int32(_a_F_assign_collations_walker_3))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L21
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	goto L1
L78:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+26)))
	if v247 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	v249 = F_exprCollation(m, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L21
	} else {
		goto L80
	}
L80:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	v253 = F_exprLocation(m, v252)
	mBase = m.M
	v854 = int32(1)
	v855 = v249
	v857 = v253
	goto L3
L81:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v854 = v260
	v855 = v261
	v857 = v259
	goto L3
L82:
	;
	v266 = F_exprLocation(m, l0)
	mBase = m.M
	v854 = base.B2i32(v262 != int32(0))
	v855 = v262
	v857 = v266
	goto L3
L83:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+76)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v645
	v653 = F_assign_collations_walker(m, v644, v15+int32(72))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L21
	} else {
		goto L179
	}
L84:
	;
	v571 = v15 + int32(48)
	v572 = m.G0
	v574 = v572 - int32(32)
	m.G0 = v574
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v576 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L21
	} else {
		goto L164
	}
L86:
	;
	v348 = m.G0
	v350 = v348 - int32(32)
	m.G0 = v350
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v352 != 0 {
		goto L105
	} else {
		goto L106
	}
L87:
	;
	v271 = v15 + int32(48)
	v272 = m.G0
	v274 = v272 - int32(32)
	m.G0 = v274
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v276 == int32(0) {
		v287 = v3
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v289 = F_assign_collations_walker(m, v288, v271)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L21
	} else {
		goto L92
	}
L89:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	if v279 != int32(1) {
		v287 = v3
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v283 = F_get_func_variadictype(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L21
	} else {
		goto L91
	}
L91:
	;
	v287 = base.B2i32(v283 == int32(0))
	goto L88
L92:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v291 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	m.G0 = v274 + int32(32)
	goto L83
L94:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if v294 <= int32(0) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v301 = int32(0)
	goto L96
L96:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v310+v301<<(uint(int32(2))%32))))
	if v287 != 0 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L93
L98:
	;
	v329 = v301 + int32(1)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if v329 < v330 {
		v301 = v329
		goto L96
	} else {
		goto L104
	}
L99:
	;
	v315 = F_assign_collations_walker(m, v314, v271)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L21
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	*(*int32)(unsafe.Add(mBase, uint32(v274)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v274)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v274)+8)) = v317
	v325 = F_assign_collations_walker(m, v314, v274+int32(8))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L21
	} else {
		goto L103
	}
L102:
	;
	goto L98
L103:
	;
	goto L98
L104:
	;
	goto L97
L105:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+12))
	v354 = v353
	goto L107
L106:
	;
	v354 = int32(0)
	goto L107
L107:
	;
	v356 = v15 + int32(48)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v357 == int32(0) {
		v370 = v352
		v371 = v3
		v372 = v3
		goto L108
	} else {
		goto L109
	}
L108:
	;
	if v370 != 0 {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v357)+12))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v357)+4))
	if v361 != int32(1) {
		v370 = v352
		v371 = v3
		v372 = v360
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v365 = F_get_func_variadictype(m, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L21
	} else {
		goto L111
	}
L111:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v370 = v367
	v371 = base.B2i32(v365 == int32(0))
	v372 = v360
	goto L108
L112:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	v375 = v373
	goto L114
L113:
	;
	v375 = int32(0)
	goto L114
L114:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v376 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	v379 = v377
	goto L117
L116:
	;
	v379 = int32(0)
	goto L117
L117:
	;
	v380 = v375 - v379
	if int32(0) < v380 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v388 = v380
	v389 = v354
	goto L121
L119:
	;
	v419 = v354
	goto L120
L120:
	;
	v425 = int32(0)
	if base.B2i32(v419 == v425)|base.B2i32(v372 == v425) != 0 {
		goto L130
	} else {
		goto L131
	}
L121:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	v396 = F_assign_collations_walker(m, v395, v356)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L21
	} else {
		goto L123
	}
L122:
	;
	v419 = v408
	goto L120
L123:
	;
	v399 = v389 + int32(4)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+12))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	if base.Ui32(v399) < base.Ui32(v402+v403<<(uint(int32(2))%32)) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v408 = v399
	goto L126
L125:
	;
	v408 = int32(0)
	goto L126
L126:
	;
	v409 = int32(1)
	if base.Ui32(v409) < base.Ui32(v388) {
		v388 = v388 - v409
		v389 = v408
		goto L121
	} else {
		goto L127
	}
L127:
	;
	goto L122
L128:
	;
	goto L83
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L21
	} else {
		goto L156
	}
L130:
	;
	m.G0 = v350 + int32(32)
	goto L128
L131:
	;
	v436 = v419
	v441 = v372
	goto L132
L132:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	*(*int32)(unsafe.Add(mBase, uint32(v350)+28)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v350)+20)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v350)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v350)+8)) = v444
	v453 = v350 + int32(8)
	v454 = F_assign_collations_walker(m, v443, v453)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L21
	} else {
		goto L134
	}
L133:
	;
	goto L130
L134:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	v457 = F_assign_collations_walker(m, v456, v453)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L21
	} else {
		goto L135
	}
L135:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v350)+16))
	if v459 == int32(2) {
		goto L129
	} else {
		goto L136
	}
L136:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v350)+12))
	if v462 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	if v371 != 0 {
		goto L144
	} else {
		goto L145
	}
L138:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	v466 = F_exprCollation(m, v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L21
	} else {
		goto L139
	}
L139:
	;
	if v466 == v462 {
		goto L137
	} else {
		goto L140
	}
L140:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	v470 = F_exprType(m, v469)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L21
	} else {
		goto L141
	}
L141:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	v473 = F_exprTypmod(m, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L21
	} else {
		goto L142
	}
L142:
	;
	v476 = F_makeRelabelType(m, v469, v470, v473, v462, int32(2))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L21
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+4)) = v476
	goto L137
L144:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v350)+20))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v350)+24))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v350)+28))
	F_merge_collation_state(m, v462, v459, v480, v481, v482, v356)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L21
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v486 = v436 + int32(4)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)+12))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v487)+4))
	v492 = v488 + v489<<(uint(int32(2))%32)
	if base.Ui32(v492) <= base.Ui32(v486) {
		goto L130
	} else {
		goto L148
	}
L147:
	;
	goto L146
L148:
	;
	v495 = v441 + int32(4)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+12))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	v503 = base.B2i32(base.Ui32(v495) < base.Ui32(v498+v499<<(uint(int32(2))%32)))
	if base.Ui32(v495) < base.Ui32(v498+v499<<(uint(int32(2))%32)) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v504 = v495
	goto L151
L150:
	;
	v504 = int32(0)
	goto L151
L151:
	;
	if base.Ui32(v486) < base.Ui32(v492) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v507 = v486
	goto L154
L153:
	;
	v507 = int32(0)
	goto L154
L154:
	;
	if base.Ui32(v495) < base.Ui32(v498+v499<<(uint(int32(2))%32)) {
		v436 = v507
		v441 = v504
		goto L132
	} else {
		goto L155
	}
L155:
	;
	goto L133
L156:
	;
	F_errcode(m, int32(17432708))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L21
	} else {
		goto L157
	}
L157:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v350)+12))
	v531 = F_get_collation_name(m, v530)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L21
	} else {
		goto L158
	}
L158:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v350)+24))
	v534 = F_get_collation_name(m, v533)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L21
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v350)+4)) = v534
	*(*int32)(unsafe.Add(mBase, uint32(v350))) = v531
	F_errmsg(m, int32(_a_F_assign_collations_walker_0), v350)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L21
	} else {
		goto L160
	}
L160:
	;
	F_errhint(m, int32(_a_F_assign_collations_walker_1), int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L21
	} else {
		goto L161
	}
L161:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v350)+8))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v350)+28))
	F_parser_errposition(m, v545, v546)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L21
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_assign_collations_walker_2), int32(1010), int32(_a_F_assign_collations_walker_4))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L21
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	v558 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+50)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v558
	F_errmsg_internal(m, int32(_a_F_assign_collations_walker_5), v15+int32(32))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L21
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_assign_collations_walker_2), int32(616), int32(_a_F_assign_collations_walker_3))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L21
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	m.G0 = v574 + int32(32)
	goto L83
L168:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	if v579 <= int32(0) {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v585 = v3
	goto L170
L170:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v576)+12))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v594+v585<<(uint(int32(2))%32))))
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+26)))
	if v599 != 0 {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	goto L167
L172:
	;
	v614 = v585 + int32(1)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	if v614 < v615 {
		v585 = v614
		goto L170
	} else {
		goto L178
	}
L173:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v571)))
	*(*int32)(unsafe.Add(mBase, uint32(v574)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v574)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v574)+8)) = v600
	v608 = F_assign_collations_walker(m, v598, v574+int32(8))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L21
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v610 = F_assign_collations_walker(m, v598, v571)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L21
	} else {
		goto L177
	}
L176:
	;
	goto L172
L177:
	;
	goto L172
L178:
	;
	goto L171
L179:
	;
	goto L5
L180:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+76)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v661
	v669 = F_assign_collations_walker(m, v660, v15+int32(72))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L21
	} else {
		goto L181
	}
L181:
	;
	goto L5
L182:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v722 = F_assign_collations_walker(m, v719, v15+int32(48))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L21
	} else {
		goto L190
	}
L183:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v671)+4))
	if v674 <= int32(0) {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v682 = int32(0)
	goto L185
L185:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v671)+12))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v690+v682<<(uint(int32(2))%32))))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v694)+4))
	v697 = v15 + int32(48)
	v698 = F_assign_collations_walker(m, v695, v697)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L21
	} else {
		goto L187
	}
L186:
	;
	goto L182
L187:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v694)+8))
	v701 = F_assign_collations_walker(m, v700, v697)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L21
	} else {
		goto L188
	}
L188:
	;
	v704 = v682 + int32(1)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v671)+4))
	if v704 < v705 {
		v682 = v704
		goto L185
	} else {
		goto L189
	}
L189:
	;
	goto L186
L190:
	;
	goto L5
L191:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+76)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v735
	v741 = F_assign_collations_walker(m, v734, v731)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L21
	} else {
		goto L192
	}
L192:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v745 = v15 + int32(48)
	v746 = F_assign_collations_walker(m, v743, v745)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L21
	} else {
		goto L193
	}
L193:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v749 = F_assign_collations_walker(m, v748, v745)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L21
	} else {
		goto L194
	}
L194:
	;
	goto L5
L195:
	;
	goto L5
L196:
	;
	F_exprSetCollation(m, l0, v791)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L21
	} else {
		goto L208
	}
L197:
	;
	v770 = F_get_typcollation(m, v768)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L21
	} else {
		goto L198
	}
L198:
	;
	if v770 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v775 = int32(0)
	v789 = v775
	v790 = v775
	v791 = v775
	v792 = int32(-1)
	goto L196
L200:
	;
	goto L201
L201:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	if v778 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v782 = F_exprLocation(m, l0)
	mBase = m.M
	v789 = int32(1)
	v790 = v770
	v791 = v770
	v792 = v782
	goto L196
L203:
	;
	goto L204
L204:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	if v778 != int32(2) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v787 = v783
	goto L207
L206:
	;
	v787 = int32(0)
	goto L207
L207:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v789 = v778
	v790 = v783
	v791 = v787
	v792 = v788
	goto L196
L208:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	if v795 == int32(2) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v803 = v801 - int32(9)
	if base.Ui32(int32(30)) < base.Ui32(v803) {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	goto L211
L211:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v828 = v826 - int32(9)
	if base.Ui32(int32(30)) < base.Ui32(v828) {
		goto L222
	} else {
		goto L223
	}
L212:
	;
	v854 = v789
	v855 = v790
	v857 = v792
	goto L3
L213:
	;
	goto L212
L214:
	;
	v807 = int32(1) << (uint(v803) % 32)
	if v807&int32(3904) == int32(0) {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v819+l0))) = int32(0)
	goto L213
L216:
	;
	if v807&int32(5) != 0 {
		v819 = int32(16)
		goto L215
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v819 = int32(24)
	goto L215
L219:
	;
	if v803 != int32(30) {
		goto L213
	} else {
		goto L220
	}
L220:
	;
	v819 = int32(12)
	goto L215
L221:
	;
	v854 = v789
	v855 = v790
	v857 = v792
	goto L3
L222:
	;
	goto L221
L223:
	;
	v832 = int32(1) << (uint(v828) % 32)
	if v832&int32(3904) == int32(0) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v844+l0))) = v823
	goto L222
L225:
	;
	if v832&int32(5) != 0 {
		v844 = int32(16)
		goto L224
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v844 = int32(24)
	goto L224
L228:
	;
	if v828 != int32(30) {
		goto L222
	} else {
		goto L229
	}
L229:
	;
	v844 = int32(12)
	goto L224
L230:
	;
	v854 = v182
	v855 = v184
	v857 = v183
	goto L3
L231:
	;
	goto L1
}
func F_assign_synchronous_commit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = l0 - int32(2)
	if base.Ui32(int32(3)) <= base.Ui32(v6) {
		v9 = int32(-1)
	} else {
		v9 = v6
	}
	*(*int32)(unsafe.Add(mBase, _c_F_assign_synchronous_commit[0])) = v9
	return
}
func F_assign_timing_clock_source(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int64
	_ = v8
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_assign_timing_clock_source[0])))
	if v4 == int32(1) {
		v8 = int64(0)
		*(*int64)(unsafe.Add(mBase, _c_F_assign_timing_clock_source[1])) = v8
		*(*int64)(unsafe.Add(mBase, _c_F_assign_timing_clock_source[2])) = v8
		*(*int32)(unsafe.Add(mBase, _c_F_assign_timing_clock_source[3])) = l0
	} else {
	}
	return
}
func F_atoi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	v5 = l0
	for {
		v10 = v5 + int32(1)
		v11 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5))))
		if base.B2i32(v11 == int32(32))|base.B2i32(base.Ui32(v11-int32(9)) < base.Ui32(int32(5))) != 0 {
			v5 = v10
			continue
		} else {
			break
		}
		break
	}
	v19 = int32(1)
	switch v11&int32(255) - int32(43) {
	case 0:
		v25 = v19
		v26 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
		v27 = v26
		v28 = v10
		v29 = v25
	default:
		v27 = v11
		v28 = v5
		v29 = v19
	case 2:
		v25 = int32(0)
		v26 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
		v27 = v26
		v28 = v10
		v29 = v25
	}
	v30 = int32(0)
	v32 = v27 - int32(48)
	if base.Ui32(v32) <= base.Ui32(int32(9)) {
		v35 = v30
		v36 = v32
		v37 = v28
		for {
			v39 = int32(10)
			v41 = v35*v39 - v36
			v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(v37)+1)))
			v46 = v42 - int32(48)
			if base.Ui32(v46) < base.Ui32(v39) {
				v35 = v41
				v36 = v46
				v37 = v37 + int32(1)
				continue
			} else {
				break
			}
			break
		}
		v49 = v41
	} else {
		v49 = v30
	}
	if v29 != 0 {
		v55 = int32(0) - v49
	} else {
		v55 = v49
	}
	return v55
}
func F_attnumAttName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 <= int32(0) {
		v12 = F_SystemAttributeDefinition(m, base.I32_extend16_s(l1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v29 = v12
			m.G0 = v7 + int32(16)
			return v29 + int32(4)
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		if v17 < l1 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg_internal(m, int32(_a_F_attnumAttName_0), v7)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_attnumAttName_1), int32(3672), int32(_a_F_attnumAttName_2))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = v16 + v17<<(uint(int32(3))%32) + l1*int32(100) - int32(72)
			m.G0 = v7 + int32(16)
			return v29 + int32(4)
		}
	}
}
func F_attnumCollationId(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 <= int32(0) {
		v24 = int32(0)
		m.G0 = v7 + int32(16)
		return v24
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v13 < l1 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg_internal(m, int32(_a_F_attnumCollationId_0), v7)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_attnumCollationId_1), int32(3712), int32(_a_F_attnumCollationId_2))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13<<(uint(int32(3))%32)+l1*int32(100))+24))
			v24 = v21
			m.G0 = v7 + int32(16)
			return v24
		}
	}
}
