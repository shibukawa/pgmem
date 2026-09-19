package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AcquireDeletionLock(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v3 - int32(1259) {
	case 0:
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_LockRelationOid(m, v6, int32(8))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	default:
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_LockDatabaseObject(m, v3, v15, int32(8))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			return
		}
	case 2:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_LockSharedObject(m, int32(1261), v11, int32(8))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			return
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
					F_errfinish(m, int32(_a_F_AllocSetRealloc_1), int32(1198), int32(_a_F_AllocSetRealloc_2))
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
			if v28 != int32(474) {
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
						F_errfinish(m, int32(_a_F_AllocSetRealloc_1), int32(1198), int32(_a_F_AllocSetRealloc_2))
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
							F_errfinish(m, int32(_a_F_AllocSetRealloc_1), int32(1198), int32(_a_F_AllocSetRealloc_2))
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
func F_AtProcExit_Buffers(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_AtProcExit_Buffers[0]))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a_F_AtProcExit_Buffers_0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(_a_F_AtProcExit_Buffers_1)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(_a_F_AtProcExit_Buffers_2)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(0)
	v20 = int32(_a_F_AtProcExit_Buffers_3)
	v22 = base.AtomicRmwOr32(m, v9, int32(24), v20)
	if v22&v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v6 + int32(32)
	return
L4:
	;
	goto L7
L5:
	;
	v38 = v22
	goto L6
L6:
	;
	v43 = int32(_a_F_AtProcExit_Buffers_4)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_AtProcExit_Buffers[1]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(8))+8))
	if v46 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	F_perform_spin_delay(m, v6+int32(8))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v38 = v34
	goto L6
L9:
	;
	return
L10:
	;
	v32 = int32(_a_F_AtProcExit_Buffers_3)
	v34 = base.AtomicRmwOr32(m, v9, int32(24), v32)
	if v34&v32 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	if v38&int32(536870912) != 0 {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtProcExit_Buffers[1])) = v61
	goto L13
L15:
	;
	if int32(999) < v44 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v44 < int32(11) {
		goto L13
	} else {
		goto L22
	}
L18:
	;
	v51 = int32(900)
	if v51 <= v44 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v54 = v51
	goto L21
L20:
	;
	v54 = v44
	goto L21
L21:
	;
	v61 = v54 + int32(100)
	goto L14
L22:
	;
	v61 = v44 - int32(1)
	goto L14
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_AtProcExit_Buffers[2]))
	if v67 == v69 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v72 = v38
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v72 & int32(-4194305)
	*(*int32)(unsafe.Add(mBase, _c_F_AtProcExit_Buffers[0])) = int32(0)
	goto L3
L26:
	;
	v71 = v38 & int32(-536870913)
	goto L28
L27:
	;
	v71 = v38
	goto L28
L28:
	;
	v72 = v71
	goto L25
}
func F_AtProcExit_Twophase(m *base.Module, l0 int32, l1 int32) {
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
	v33 = v22 + int32(640)
	goto L4
L6:
	;
	v18 = v15 * int32(640)
	v19 = v11 + v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
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
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+684))
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
func F_accumArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
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
				v30 = F_MemoryContextAlloc(m, v17, int32(256))
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
								v62 = v55 << (uint(int32(3)) % 32)
								if base.Ui32(int32(1073741824)) <= base.Ui32(v62) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(261))
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(1073741823)
											F_errmsg(m, int32(_a_F_accumArrayResult_4), v9)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_accumArrayResult_5), int32(_a_F_accumArrayResult_6), int32(_a_F_accumArrayResult_0))
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
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
												v84 = l1
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
												*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
												m.G0 = v9 + int32(16)
												return v48
											} else {
												v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
												if v75 != 0 {
													v84 = l1
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
													*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
													v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
													*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
													v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
													*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
													*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
													m.G0 = v9 + int32(16)
													return v48
												} else {
													v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+24)))
													if v76 == int32(-1) {
														v79 = F_pg_detoast_datum_copy(m, l1)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return int32(0)
														} else {
															v84 = v79
															v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
															v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
															v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
															*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
															m.G0 = v9 + int32(16)
															return v48
														}
													} else {
														v82 = F_datumCopy(m, l1, int32(0), v76)
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
															return int32(0)
														} else {
															v84 = v82
															v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
															v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
															v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
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
									v84 = l1
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
									*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
									m.G0 = v9 + int32(16)
									return v48
								} else {
									v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
									if v75 != 0 {
										v84 = l1
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
										*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
										m.G0 = v9 + int32(16)
										return v48
									} else {
										v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+24)))
										if v76 == int32(-1) {
											v79 = F_pg_detoast_datum_copy(m, l1)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												v84 = v79
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
												*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
												m.G0 = v9 + int32(16)
												return v48
											}
										} else {
											v82 = F_datumCopy(m, l1, int32(0), v76)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												v84 = v82
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
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
			v62 = v55 << (uint(int32(3)) % 32)
			if base.Ui32(int32(1073741824)) <= base.Ui32(v62) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(1073741823)
						F_errmsg(m, int32(_a_F_accumArrayResult_4), v9)
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_accumArrayResult_5), int32(_a_F_accumArrayResult_6), int32(_a_F_accumArrayResult_0))
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
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
							v84 = l1
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
							m.G0 = v9 + int32(16)
							return v48
						} else {
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
							if v75 != 0 {
								v84 = l1
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
								*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
								m.G0 = v9 + int32(16)
								return v48
							} else {
								v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+24)))
								if v76 == int32(-1) {
									v79 = F_pg_detoast_datum_copy(m, l1)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										v84 = v79
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
										*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
										m.G0 = v9 + int32(16)
										return v48
									}
								} else {
									v82 = F_datumCopy(m, l1, int32(0), v76)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										v84 = v82
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
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
				v84 = l1
				v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
				v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
				v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
				*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
				m.G0 = v9 + int32(16)
				return v48
			} else {
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
				if v75 != 0 {
					v84 = l1
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
					v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
					*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
					m.G0 = v9 + int32(16)
					return v48
				} else {
					v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+24)))
					if v76 == int32(-1) {
						v79 = F_pg_detoast_datum_copy(m, l1)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v84 = v79
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_accumArrayResult[0])) = v51
							m.G0 = v9 + int32(16)
							return v48
						}
					} else {
						v82 = F_datumCopy(m, l1, int32(0), v76)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v84 = v82
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
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
				F_errfinish(m, int32(_a_F_aclcopy_1), int32(433), int32(_a_F_aclcopy_2))
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
func F_aclinsert(m *base.Module, l0 int32) int32 {
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
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(_a_F_aclinsert_0), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_aclinsert_1), int32(1597), int32(_a_F_aclinsert_2))
				v19 = m.ExcPending
				if v19 != 0 {
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
						F_pg_qsort(m, v21, v65, int32(4), int32(471))
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
						F_pg_qsort(m, v21, int32(0), int32(4), int32(471))
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 float64
	_ = v101
	var v104 float64
	_ = v104
	var v108 float64
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
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
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 float64
	_ = v156
	var v157 float64
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v187 float64
	_ = v187
	var v188 float64
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v205 float64
	_ = v205
	var v217 float64
	_ = v217
	var v238 float64
	_ = v238
	var v242 float64
	_ = v242
	var v244 float64
	_ = v244
	var v251 float64
	_ = v251
	var v252 float64
	_ = v252
	var v255 float64
	_ = v255
	var v263 float64
	_ = v263
	var v264 float64
	_ = v264
	var v266 float64
	_ = v266
	var v269 float64
	_ = v269
	var v271 float64
	_ = v271
	var v272 float64
	_ = v272
	var v273 float64
	_ = v273
	var v275 float64
	_ = v275
	var v276 float64
	_ = v276
	var v278 int32
	_ = v278
	var v279 float64
	_ = v279
	var v283 float64
	_ = v283
	var v295 float64
	_ = v295
	var v300 float64
	_ = v300
	var v301 float64
	_ = v301
	var v302 float64
	_ = v302
	var v306 float64
	_ = v306
	var v308 float64
	_ = v308
	var v310 float64
	_ = v310
	var v313 float64
	_ = v313
	var v316 float64
	_ = v316
	var v322 float64
	_ = v322
	var v323 int32
	_ = v323
	var v324 float64
	_ = v324
	var v327 float64
	_ = v327
	var v331 float64
	_ = v331
	var v332 float64
	_ = v332
	var v336 float64
	_ = v336
	var v344 float64
	_ = v344
	var v345 float64
	_ = v345
	var v348 float64
	_ = v348
	var v358 float64
	_ = v358
	var v380 float64
	_ = v380
	var v383 float64
	_ = v383
	var v385 float64
	_ = v385
	var v386 float64
	_ = v386
	var v389 float64
	_ = v389
	var v397 float64
	_ = v397
	var v417 float64
	_ = v417
	var v425 float64
	_ = v425
	var v431 float64
	_ = v431
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v452 float64
	_ = v452
	var v454 float64
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v473 int32
	_ = v473
	var v480 float64
	_ = v480
	var v481 float64
	_ = v481
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v536 int32
	_ = v536
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 float64
	_ = v564
	var v566 float64
	_ = v566
	var v567 float64
	_ = v567
	var v569 float64
	_ = v569
	var v571 float64
	_ = v571
	var v574 float64
	_ = v574
	var v580 float64
	_ = v580
	var v583 float64
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 float64
	_ = v590
	var v591 float64
	_ = v591
	var v593 float64
	_ = v593
	var v597 int32
	_ = v597
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	v7 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(128)
	m.G0 = v22
	v24 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+120)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v22)+112)) = v24
	v29 = F_RelationGetNumberOfBlocksInFork(m, l0, v7)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = F_GetOldestNonRemovableTransactionId(m, l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = v22 + int32(80)
	v39 = Fn13986(m, int64(32))
	mBase = m.M
	goto L4
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v29
	F_pg_prng_seed(m, v22+int32(96), base.I64_extend_i32_u(v39))
	mBase = m.M
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if base.Ui32(v48) < base.Ui32(v49) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[0]))
	if v55 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v51 = v48
	goto L8
L7:
	;
	v51 = v49
	goto L8
L8:
	;
	goto L5
L9:
	;
	v93 = v22 + int32(64)
	v94 = F_pg_prng_uint32(m)
	mBase = m.M
	F_pg_prng_seed(m, v93, base.I64_extend_i32_u(v94))
	mBase = m.M
	goto L14
L10:
	;
	goto L9
L11:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_acquire_sample_rows[1])))
	if v59&int32(1) == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v64 = int32(_a_F_acquire_sample_rows_0)
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2]))
	v67 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2])) = v66 + v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v70 + v67
	*(*int64)(unsafe.Add(mBase, uint32(v55+int32(8))+232)) = base.I64_extend_i32_u(v51)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v78 + v67
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2])) = v84 - v67
	goto L10
L13:
	;
	v110 = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v117 = m.T0[v116].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v110, v110, v110, v110, int32(32))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L17
	}
L14:
	;
	v101 = F_pg_prng_double(m, v93)
	mBase = m.M
	if base.F64_eq(v101, float64(0)) != 0 {
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v104 = F_log(m, v101)
	mBase = m.M
	v108 = F_exp(m, base.F64_div(base.F64_neg(v104), base.F64_convert_i32_s(l3)))
	mBase = m.M
	*(*float64)(unsafe.Add(mBase, uint32(v22+int32(56)))) = v108
	goto L13
L16:
	;
	goto L15
L17:
	;
	v120 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[3]))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v126 = int32(0)
	v129 = F_read_stream_begin_relation(m, int32(9), v124, v125, v126, int32(502), v37, v126)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+188))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+132))
	v134 = m.T0[v133].(func(*base.Module, int32, int32) int32)(m, v117, v129)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v134 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v149 = v7
	v155 = v7
	v156 = float64(-1)
	v157 = float64(0)
	goto L24
L22:
	;
	v536 = v7
	goto L23
L23:
	;
	F_read_stream_end(m, v129)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L93
	}
L24:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v536 = v473
	goto L23
L26:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+188))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+136))
	v169 = m.T0[v168].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v117, v33, v22+int32(120), v22+int32(112), v120)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v169 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v180 = v149
	v187 = v156
	v188 = v157
	goto L31
L29:
	;
	v473 = v149
	v480 = v156
	v481 = v157
	goto L30
L30:
	;
	v485 = v155 + int32(1)
	v489 = *(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[0]))
	if v489 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L31:
	;
	if v180 < l3 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v473 = v450
	v480 = v452
	v481 = v454
	goto L30
L33:
	;
	v454 = base.F64_add(v188, float64(1))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+188))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)+136))
	v462 = m.T0[v461].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v117, v33, v22+int32(120), v22+int32(112), v120)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L85
	}
L34:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+44))
	v196 = m.T0[v195].(func(*base.Module, int32) int32)(m, v120)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if base.F64_lt(v187, float64(0)) != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2+v180<<(uint(int32(2))%32)))) = v196
	v450 = v180 + int32(1)
	v452 = v187
	goto L33
L38:
	;
	v204 = v22 + int32(56)
	v205 = float64(0)
	v217 = base.F64_convert_i32_s(l3)
	if base.F64_ge(base.F64_mul(v217, float64(22)), v188) != 0 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	v425 = v187
	goto L40
L40:
	;
	if base.F64_le(v425, float64(0)) != 0 {
		goto L76
	} else {
		goto L77
	}
L41:
	;
	v425 = v417
	goto L40
L42:
	;
	goto L41
L43:
	;
	goto L46
L44:
	;
	goto L45
L45:
	;
	v271 = float64(1)
	v272 = base.F64_add(v188, v271)
	v273 = base.F64_sub(v188, v217)
	v275 = base.F64_add(v273, v271)
	v276 = base.F64_div(v272, v275)
	v278 = v22 + int32(64)
	v279 = *(*float64)(unsafe.Add(mBase, uint32(v204)))
	v283 = v279
	goto L53
L46:
	;
	v238 = F_pg_prng_double(m, v22+int32(64))
	mBase = m.M
	if base.F64_eq(v238, float64(0)) != 0 {
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v242 = base.F64_add(v188, float64(1))
	v244 = base.F64_div(base.F64_sub(v242, v217), v242)
	if base.F64_gt(v244, v238) == int32(0) {
		v417 = v205
		goto L42
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v251 = v242
	v252 = v244
	v255 = v205
	goto L50
L50:
	;
	v263 = float64(1)
	v264 = base.F64_add(v255, v263)
	v266 = base.F64_add(v251, v263)
	v269 = base.F64_mul(v252, base.F64_div(base.F64_sub(v266, v217), v266))
	if base.F64_gt(v269, v238) != 0 {
		v251 = v266
		v252 = v269
		v255 = v264
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v417 = v264
	goto L42
L52:
	;
	goto L51
L53:
	;
	v295 = F_pg_prng_double(m, v278)
	mBase = m.M
	if base.F64_eq(v295, float64(0)) != 0 {
		goto L53
	} else {
		goto L55
	}
L54:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v204))) = v397
	v417 = v301
	goto L42
L55:
	;
	v300 = base.F64_mul(v188, base.F64_add(v283, float64(-1)))
	v301 = base.F64_floor(v300)
	v302 = base.F64_add(v275, v301)
	v306 = base.F64_add(v188, v300)
	v308 = F_log(m, base.F64_div(base.F64_mul(v302, base.F64_mul(v276, base.F64_mul(v276, v295))), v306))
	mBase = m.M
	v310 = F_exp(m, base.F64_div(v308, v217))
	mBase = m.M
	v313 = base.F64_div(base.F64_mul(v275, base.F64_div(v306, v302)), v188)
	if base.F64_le(v310, v313) != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L54
L57:
	;
	v397 = base.F64_div(v313, v310)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v316 = base.F64_add(v188, v301)
	v322 = base.F64_div(base.F64_mul(base.F64_add(v316, float64(1)), base.F64_div(base.F64_mul(v272, v295), v275)), v306)
	v323 = base.F64_lt(v217, v301)
	if v323 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v324 = v302
	goto L62
L61:
	;
	v324 = v272
	goto L62
L62:
	;
	if base.F64_le(v324, v316) != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	if v323 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v358 = v322
	goto L65
L65:
	;
	goto L72
L66:
	;
	v327 = v188
	goto L68
L67:
	;
	v327 = base.F64_add(v273, v301)
	goto L68
L68:
	;
	v331 = v316
	v332 = v327
	v336 = v322
	goto L69
L69:
	;
	v344 = base.F64_mul(v336, base.F64_div(v331, v332))
	v345 = float64(-1)
	v348 = base.F64_add(v331, v345)
	if base.F64_ge(v348, v324) != 0 {
		v331 = v348
		v332 = base.F64_add(v332, v345)
		v336 = v344
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v358 = v344
	goto L65
L71:
	;
	goto L70
L72:
	;
	v380 = F_pg_prng_double(m, v278)
	mBase = m.M
	if base.F64_eq(v380, float64(0)) != 0 {
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v383 = F_log(m, v358)
	mBase = m.M
	v385 = F_exp(m, base.F64_div(v383, v217))
	mBase = m.M
	v386 = F_log(m, v380)
	mBase = m.M
	v389 = F_exp(m, base.F64_div(base.F64_neg(v386), v217))
	mBase = m.M
	if base.F64_le(v385, base.F64_div(v306, v188)) == int32(0) {
		v283 = v389
		goto L53
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	v397 = v389
	goto L56
L76:
	;
	goto L80
L77:
	;
	goto L78
L78:
	;
	v450 = v180
	v452 = base.F64_add(v425, float64(-1))
	goto L33
L79:
	;
	v438 = l2 + base.I32_trunc_sat_f64_s(base.F64_mul(v431, base.F64_convert_i32_s(l3)))<<(uint(int32(2))%32)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)))
	F_pfree(m, v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L83
	}
L80:
	;
	v431 = F_pg_prng_double(m, v22-int32(-64))
	mBase = m.M
	if base.F64_eq(v431, float64(0)) != 0 {
		goto L80
	} else {
		goto L82
	}
L81:
	;
	goto L79
L82:
	;
	goto L81
L83:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)+44))
	v444 = m.T0[v443].(func(*base.Module, int32) int32)(m, v120)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v438))) = v444
	goto L78
L85:
	;
	if v462 != 0 {
		v180 = v450
		v187 = v452
		v188 = v454
		goto L31
	} else {
		goto L86
	}
L86:
	;
	goto L32
L87:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)+188))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+132))
	v525 = m.T0[v524].(func(*base.Module, int32, int32) int32)(m, v117, v129)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L91
	}
L88:
	;
	goto L87
L89:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_acquire_sample_rows[1])))
	if v493&int32(1) == int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v498 = int32(_a_F_acquire_sample_rows_0)
	v500 = *(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2]))
	v501 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2])) = v500 + v501
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	*(*int32)(unsafe.Add(mBase, uint32(v489))) = v504 + v501
	*(*int64)(unsafe.Add(mBase, uint32(v489+int32(16))+232)) = base.I64_extend_i32_u(v485)
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	*(*int32)(unsafe.Add(mBase, uint32(v489))) = v512 + v501
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_acquire_sample_rows[2])) = v518 - v501
	goto L88
L91:
	;
	if v525 != 0 {
		v149 = v473
		v155 = v485
		v156 = v480
		v157 = v481
		goto L24
	} else {
		goto L92
	}
L92:
	;
	goto L25
L93:
	;
	F_ExecDropSingleTupleTableSlot(m, v120)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v550)+188))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+12))
	m.T0[v552].(func(*base.Module, int32))(m, v117)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	if l3 == v536 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_qsort_interruptible(m, l2, l3, int32(4), int32(503), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	if v561 <= int32(0) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L98
L100:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v583
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = v580
	v587 = F_errstart(m, l1, int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L104
	}
L101:
	;
	v564 = float64(0)
	v580 = v564
	v583 = v564
	goto L100
L102:
	;
	goto L103
L103:
	;
	v566 = *(*float64)(unsafe.Add(mBase, uint32(v22)+112))
	v567 = base.F64_convert_i32_u(v561)
	v569 = base.F64_convert_i32_u(v29)
	v571 = float64(0.5)
	v574 = *(*float64)(unsafe.Add(mBase, uint32(v22)+120))
	v580 = base.F64_floor(base.F64_add(base.F64_mul(base.F64_div(v566, v567), v569), v571))
	v583 = base.F64_floor(base.F64_add(base.F64_mul(base.F64_div(v574, v567), v569), v571))
	goto L100
L104:
	;
	if v587 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v590 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	v591 = *(*float64)(unsafe.Add(mBase, uint32(v22)+120))
	*(*float64)(unsafe.Add(mBase, uint32(v22)+16)) = v591
	v593 = *(*float64)(unsafe.Add(mBase, uint32(v22)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v22)+24)) = v593
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v536
	*(*float64)(unsafe.Add(mBase, uint32(v22)+40)) = v590
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v589 + int32(4)
	F_errmsg(m, int32(_a_F_acquire_sample_rows_1), v22)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	m.G0 = v22 + int32(128)
	return v536
L108:
	;
	F_errfinish(m, int32(_a_F_acquire_sample_rows_2), int32(1352), int32(_a_F_acquire_sample_rows_3))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	goto L107
}
func F_addCompoundAffixFlagValue(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	v7 = m.G0
	v9 = v7 - int32(1024)
	m.G0 = v9
	v12 = l1
	goto L2
L1:
	;
	v49 = v12
	v51 = v17
	v52 = v9
	goto L15
L2:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if base.B2i32(base.Ui32(v17-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v17 == int32(32)) == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L9
	} else {
		goto L11
	}
L4:
	;
	goto L3
L5:
	;
	if v17 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v29 = F_pg_mblen_cstr(m, v12)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L1
L9:
	;
	return
L10:
	;
	v12 = v29 + v12
	goto L2
L11:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	F_errmsg(m, int32(_a_F_addCompoundAffixFlagValue_0), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_addCompoundAffixFlagValue_1), int32(1076), int32(_a_F_addCompoundAffixFlagValue_2))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L15:
	;
	switch v51 {
	case 0, 9, 10, 11, 12, 13, 32:
		goto L17
	default:
		goto L18
	}
L16:
	;
	v60 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v60)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v63 < v62 {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	goto L16
L18:
	;
	v54 = F_pg_mblen_cstr(m, v49)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	if v54 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	base.MemoryCopy(m, v52, v49, v54)
	goto L22
L21:
	;
	goto L22
L22:
	;
	v58 = v49 + v54
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v49 = v58
	v51 = v59
	v52 = v54 + v52
	goto L15
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_setCompoundAffixFlagValue(m, l0, v82+v83*int32(12), v9, l2)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L33
	}
L24:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v82 = v65
	goto L23
L25:
	;
	goto L26
L26:
	;
	if v62 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v80
	v82 = v80
	goto L23
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62 << (uint(int32(1)) % 32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v72 = F_repalloc(m, v69, v62*int32(24))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L9
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(10)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v78 = F_MemoryContextAlloc(m, v76, int32(120))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L32
	}
L31:
	;
	v80 = v72
	goto L27
L32:
	;
	v80 = v78
	goto L27
L33:
	;
	v89 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)) = uint8(v89)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v91 + v89
	m.G0 = v9 + int32(1024)
	return
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
	F_errfinish(m, int32(_a_F_addFamilyMember_2), int32(1442), int32(_a_F_addFamilyMember_3))
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
	F_errfinish(m, int32(_a_F_addFamilyMember_2), int32(1435), int32(_a_F_addFamilyMember_3))
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
	v150 = F_palloc0(m, int32(140))
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
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
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
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	if v16 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if int32(0) < v20 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = v3
	goto L6
L4:
	;
	goto L5
L5:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+52))
	if v63 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v27<<(uint(int32(2))%32))))
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v49 = v27 + int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v49 < v50 {
		v27 = v49
		goto L6
	} else {
		goto L15
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	F_add_rte_to_flat_rtable(m, v19, v45, v37)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	switch v40 {
	case 0:
		goto L9
	case 1:
		goto L11
	default:
		goto L8
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v41 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	return
L14:
	;
	goto L8
L15:
	;
	goto L7
L16:
	;
	v66 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v67 <= v66 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v73 = int32(1)
	v75 = v66
	goto L18
L18:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+v75<<(uint(int32(2))%32))))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	if v86 != int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L1
L20:
	;
	v149 = int32(1)
	v152 = v75 + v149
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v152 < v153 {
		v73 = v73 + v149
		v75 = v152
		goto L18
	} else {
		goto L45
	}
L21:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+20)))
	if v89 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(v90) <= base.Ui32(v73) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v73<<(uint(int32(2))%32))))
	if v96 == int32(0) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+140))
	if v99 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v19
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v85)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v103
	v109 = F_query_tree_walker_impl(m, v103, int32(834), v13+int32(8), int32(16))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L13
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if l1 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L20
L29:
	;
	v142 = v99
	goto L31
L30:
	;
	v113 = F_fetch_upper_rel(m, v99, int32(7), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L13
	} else {
		goto L32
	}
L31:
	;
	F_add_rtes_to_flat_rtable(m, v142, int32(1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L44
	}
L32:
	;
	v115 = int32(0)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+32))
	if v117 == v115 {
		v138 = v115
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v138 == int32(0) {
		goto L20
	} else {
		goto L43
	}
L34:
	;
	goto L33
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v121 = v120
	goto L36
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if base.Ui32(int32(2)) <= base.Ui32(v125-int32(301)) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v138 = int32(1)
	goto L34
L38:
	;
	if v125 != int32(290) {
		v138 = v115
		goto L34
	} else {
		goto L41
	}
L39:
	;
	v121 = v124 + int32(72)
	goto L36
L40:
	;
	goto L37
L41:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v124)+72))
	if v132 != 0 {
		v138 = v115
		goto L34
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v96)+140))
	v142 = v141
	goto L31
L44:
	;
	goto L20
L45:
	;
	goto L19
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
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
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
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+76))
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
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v63)+120))
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
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
	v21 = base.I32_div_s(int32(0), int32(32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v22 <= v21 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v25 = l0 + int32(8)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v21<<(uint(int32(2))%32))))
	v32 = v29 & int32(-1)
	if v32 != 0 {
		v54 = v32
		v55 = v21
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v34 = v21 + int32(1)
	if v34 == v22 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v37 = v34
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
	v185 = base.I32_div_s(v183, int32(32))
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
	F_errfinish(m, int32(_a_F_adjust_view_column_set_2), int32(3099), int32(_a_F_adjust_view_column_set_3))
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
func F_akeys(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_akeys(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_anyarray_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_anyarray_in_0), int32(154), int32(_a_F_anyarray_in_1), int32(_a_F_anyarray_in_2), int32(_a_F_anyarray_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anycompatiblearray_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_anycompatiblearray_in_0), int32(174), int32(_a_F_anycompatiblearray_in_1), int32(_a_F_anycompatiblearray_in_2), int32(_a_F_anycompatiblearray_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anycompatiblenonarray_out(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_anycompatiblenonarray_out_0), int32(377), int32(_a_F_anycompatiblenonarray_out_1), int32(_a_F_anycompatiblenonarray_out_2), int32(_a_F_anycompatiblenonarray_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anycompatiblerange_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_anycompatiblerange_in_0), int32(220), int32(_a_F_anycompatiblerange_in_1), int32(_a_F_anycompatiblerange_in_2), int32(_a_F_anycompatiblerange_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anymultirange_out(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_multirange_out(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_anynonarray_out(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_anynonarray_out_0), int32(375), int32(_a_F_anynonarray_out_1), int32(_a_F_anynonarray_out_2), int32(_a_F_anynonarray_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anytimestamp_typmod_check(m *base.Module, l0 int32, l1 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn13867(m, l0, l1, int32(_a_F_anytimestamp_typmod_check_0), int32(131), int32(_a_F_anytimestamp_typmod_check_1), int32(_a_F_anytimestamp_typmod_check_2), int32(138), int32(_a_F_anytimestamp_typmod_check_3))
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v249 int32
	_ = v249
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
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
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
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
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v451 int32
	_ = v451
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	v7 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	if v19 == v7 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l4 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v32 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v34 == v32 {
		v55 = v32
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+240))
	if v22 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+236))
	if v25 <= int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
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
	v63 = int32(0)
	goto L1
L10:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L20
	}
L11:
	;
	goto L10
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v38 = v37
	goto L13
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if base.Ui32(int32(2)) <= base.Ui32(v42-int32(301)) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v55 = int32(1)
	goto L11
L15:
	;
	if v42 != int32(290) {
		v55 = v32
		goto L11
	} else {
		goto L18
	}
L16:
	;
	v38 = v41 + int32(72)
	goto L13
L17:
	;
	goto L14
L18:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)+72))
	if v49 != 0 {
		v55 = v32
		goto L11
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	if v55 != 0 {
		v63 = int32(0)
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(0)
	v63 = int32(1)
	goto L1
L22:
	;
	F_generate_useful_gather_paths(m, l0, l1, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L8
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v63 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v69)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v69
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(0)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v77 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v127 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L30:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v80 <= int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v90 = int32(0)
	goto L32
L32:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v101 = v98 + v90<<(uint(int32(2))%32)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if l5 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L29
L34:
	;
	v110 = v90 + int32(1)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v110 < v111 {
		v90 = v110
		goto L32
	} else {
		goto L39
	}
L35:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v104
	goto L34
L36:
	;
	goto L37
L37:
	;
	v106 = F_create_projection_path(m, l0, l1, v102, v76)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v106
	goto L34
L39:
	;
	goto L33
L40:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+38)))
	if v180 != 0 {
		goto L51
	} else {
		goto L52
	}
L41:
	;
	v130 = int32(0)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v131 <= v130 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v140 = v130
	goto L43
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v151 = v148 + v140<<(uint(int32(2))%32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	if l5 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L40
L45:
	;
	v162 = v140 + int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v162 < v163 {
		v140 = v162
		goto L43
	} else {
		goto L50
	}
L46:
	;
	v155 = F_create_projection_path(m, l0, l1, v152, v76)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L8
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v158)+8)) = v159
	goto L45
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v155
	goto L45
L50:
	;
	goto L44
L51:
	;
	F_adjust_paths_for_srfs(m, l0, l1, l2, l3)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L8
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v183+v184<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v190
	if v63 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L53
L55:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	if v192 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	goto L57
L57:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v484 != int32(1) {
		goto L123
	} else {
		goto L124
	}
L58:
	;
	if int32(0) <= v249 {
		goto L69
	} else {
		goto L70
	}
L59:
	;
	v249 = base.I32_ctz(v235) | v236<<(uint(int32(5))%32)
	goto L58
L60:
	;
	v249 = int32(-2)
	goto L58
L61:
	;
	v202 = base.I32_div_s(int32(0), int32(32))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	if v203 <= v202 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v206 = v192 + int32(8)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v206+v202<<(uint(int32(2))%32))))
	v213 = v210 & int32(-1)
	if v213 != 0 {
		v235 = v213
		v236 = v202
		goto L59
	} else {
		goto L63
	}
L63:
	;
	v215 = v202 + int32(1)
	if v215 == v203 {
		goto L60
	} else {
		goto L64
	}
L64:
	;
	v218 = v215
	goto L65
L65:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v206+v218<<(uint(int32(2))%32))))
	if v225 != 0 {
		v235 = v225
		v236 = v218
		goto L59
	} else {
		goto L67
	}
L66:
	;
	goto L60
L67:
	;
	v227 = v218 + int32(1)
	if v227 != v203 {
		v218 = v227
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v260 = v249
	v265 = v7
	goto L72
L70:
	;
	v467 = v7
	goto L71
L71:
	;
	F_add_paths_to_append_rel(m, l0, l1, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L8
	} else {
		goto L122
	}
L72:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v266+v260<<(uint(int32(2))%32))))
	v271 = int32(0)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v270)+32))
	if v273 == v271 {
		v294 = v271
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v467 = v394
	goto L71
L74:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	if v395 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L75:
	;
	if v294 != 0 {
		v394 = v265
		goto L74
	} else {
		goto L85
	}
L76:
	;
	goto L75
L77:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	v277 = v276
	goto L78
L78:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	if base.Ui32(int32(2)) <= base.Ui32(v281-int32(301)) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v294 = int32(1)
	goto L76
L80:
	;
	if v281 != int32(290) {
		v294 = v271
		goto L76
	} else {
		goto L83
	}
L81:
	;
	v277 = v280 + int32(72)
	goto L78
L82:
	;
	goto L79
L83:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v280)+72))
	if v288 != 0 {
		v294 = v271
		goto L76
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v270)+8))
	v298 = F_find_appinfos_by_relids(m, l0, v295, v17+int32(12))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	v300 = int32(0)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v300 < v302 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v312 = v300
	v315 = v300
	goto L90
L88:
	;
	v347 = v300
	goto L89
L89:
	;
	F_pfree(m, v298)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L8
	} else {
		goto L96
	}
L90:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319+v312<<(uint(int32(2))%32))))
	v324 = F_copy_pathtarget(m, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L8
	} else {
		goto L92
	}
L91:
	;
	v347 = v331
	goto L89
L92:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v328 = F_adjust_appendrel_attrs(m, l0, v326, v327, v298)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v324)+4)) = v328
	v331 = F_lappend(m, v315, v324)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	v334 = v312 + int32(1)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v334 < v335 {
		v312 = v334
		v315 = v331
		goto L90
	} else {
		goto L95
	}
L95:
	;
	goto L91
L96:
	;
	F_apply_scanjoin_target_to_paths(m, l0, v270, v347, l3, l4, l5)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L8
	} else {
		goto L97
	}
L97:
	;
	v355 = int32(0)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v270)+32))
	if v357 == v355 {
		v378 = v355
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v378 != 0 {
		v394 = v265
		goto L74
	} else {
		goto L108
	}
L99:
	;
	goto L98
L100:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v357)+12))
	v361 = v360
	goto L101
L101:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	if base.Ui32(int32(2)) <= base.Ui32(v365-int32(301)) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v378 = int32(1)
	goto L99
L103:
	;
	if v365 != int32(290) {
		v378 = v355
		goto L99
	} else {
		goto L106
	}
L104:
	;
	v361 = v364 + int32(72)
	goto L101
L105:
	;
	goto L102
L106:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v364)+72))
	if v372 != 0 {
		v378 = v355
		goto L99
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v379 = F_lappend(m, v265, v270)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	v394 = v379
	goto L74
L110:
	;
	if int32(0) <= v451 {
		v260 = v451
		v265 = v394
		goto L72
	} else {
		goto L121
	}
L111:
	;
	v451 = base.I32_ctz(v437) | v438<<(uint(int32(5))%32)
	goto L110
L112:
	;
	v451 = int32(-2)
	goto L110
L113:
	;
	v402 = v260 + int32(1)
	v404 = base.I32_div_s(v402, int32(32))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	if v405 <= v404 {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v408 = v395 + int32(8)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v408+v404<<(uint(int32(2))%32))))
	v415 = v412 & (int32(-1) << (uint(v402) % 32))
	if v415 != 0 {
		v437 = v415
		v438 = v404
		goto L111
	} else {
		goto L115
	}
L115:
	;
	v417 = v404 + int32(1)
	if v417 == v405 {
		goto L112
	} else {
		goto L116
	}
L116:
	;
	v420 = v417
	goto L117
L117:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v408+v420<<(uint(int32(2))%32))))
	if v427 != 0 {
		v437 = v427
		v438 = v420
		goto L111
	} else {
		goto L119
	}
L118:
	;
	goto L112
L119:
	;
	v429 = v420 + int32(1)
	if v429 != v405 {
		v420 = v429
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	goto L73
L122:
	;
	goto L57
L123:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L8
	} else {
		goto L130
	}
L124:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(1)<<(uint(v487)%32)&int32(44) != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v495 = base.B2i32(base.Ui32(v487) <= base.Ui32(int32(5)))
	goto L127
L126:
	;
	v495 = int32(0)
	goto L127
L127:
	;
	if v495 != 0 {
		goto L123
	} else {
		goto L128
	}
L128:
	;
	F_generate_useful_gather_paths(m, l0, l1, int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L8
	} else {
		goto L129
	}
L129:
	;
	goto L123
L130:
	;
	m.G0 = v17 + int32(16)
	return
}
func F_apw_detach_shmem(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_apw_detach_shmem[0]))
	v6 = F_LWLockAcquire(m, v4, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_apw_detach_shmem[1]))
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_apw_detach_shmem[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
		if v9 == v12 {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(-1)
		} else {
		}
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		if v9 == v16 {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(-1)
		} else {
		}
		F_LWLockRelease(m, v11)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			return
		}
	}
}
func F_apw_init_state(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = F_LWLockNewTrancheId(m)
	mBase = m.M
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v2)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1073741824)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-1)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-1)
		return
	}
}
func F_arraycontsel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 float64
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 float64
	_ = v88
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
	var v185 int32
	_ = v185
	var v190 float64
	_ = v190
	var v198 float64
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 float64
	_ = v203
	var v218 float64
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	v9 = float64(0)
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v25 = F_get_restriction_variable(m, v16, v17, v18, v13+int32(8), v13+int32(4), v13+int32(3))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return int32(0)
	} else {
		if v25 == int32(0) {
			if v15 == int32(2750) {
				v35 = float64(0.01)
			} else {
				v35 = float64(0.005)
			}
			v218 = v35
			v220 = F_Float8GetDatum(m, v218)
			mBase = m.M
			v221 = m.ExcPending
			if v221 != 0 {
				return int32(0)
			} else {
				m.G0 = v13 + int32(112)
				return v220
			}
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
			if v37 != int32(7) {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
				if v40 != 0 {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
					m.T0[v41].(func(*base.Module, int32))(m, v40)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						if v15 == int32(2750) {
							v48 = float64(0.01)
						} else {
							v48 = float64(0.005)
						}
						v218 = v48
						v220 = F_Float8GetDatum(m, v218)
						mBase = m.M
						v221 = m.ExcPending
						if v221 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(112)
							return v220
						}
					}
				} else {
					if v15 == int32(2750) {
						v48 = float64(0.01)
					} else {
						v48 = float64(0.005)
					}
					v218 = v48
					v220 = F_Float8GetDatum(m, v218)
					mBase = m.M
					v221 = m.ExcPending
					if v221 != 0 {
						return int32(0)
					} else {
						m.G0 = v13 + int32(112)
						return v220
					}
				}
			} else {
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)))
				if v49 == int32(1) {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
					if v52 == int32(0) {
						v218 = v9
						v220 = F_Float8GetDatum(m, v218)
						mBase = m.M
						v221 = m.ExcPending
						if v221 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(112)
							return v220
						}
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
						m.T0[v55].(func(*base.Module, int32))(m, v52)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v218 = v9
							v220 = F_Float8GetDatum(m, v218)
							mBase = m.M
							v221 = m.ExcPending
							if v221 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(112)
								return v220
							}
						}
					}
				} else {
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
					if v58 != 0 {
						v66 = v15
					} else {
						if v15 == int32(2751) {
							v66 = int32(2752)
						} else {
							if v15 == int32(2752) {
								v65 = int32(2751)
							} else {
								v65 = v15
							}
							v66 = v65
						}
					}
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
					v68 = F_get_base_element_type(m, v67)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						if v68 == int32(0) {
							if v66 == int32(2750) {
								v190 = float64(0.01)
							} else {
								v190 = float64(0.005)
							}
							v198 = v190
							v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
							if v199 != 0 {
								v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
								m.T0[v200].(func(*base.Module, int32))(m, v199)
								mBase = m.M
								v202 = m.ExcPending
								if v202 != 0 {
									return int32(0)
								} else {
									v203 = float64(0)
									if base.F64_lt(v198, v203) != 0 {
										v218 = v203
									} else {
										if base.F64_gt(v198, float64(1)) == int32(0) {
											v218 = v198
										} else {
											v218 = float64(1)
										}
									}
									v220 = F_Float8GetDatum(m, v218)
									mBase = m.M
									v221 = m.ExcPending
									if v221 != 0 {
										return int32(0)
									} else {
										m.G0 = v13 + int32(112)
										return v220
									}
								}
							} else {
								v203 = float64(0)
								if base.F64_lt(v198, v203) != 0 {
									v218 = v203
								} else {
									if base.F64_gt(v198, float64(1)) == int32(0) {
										v218 = v198
									} else {
										v218 = float64(1)
									}
								}
								v220 = F_Float8GetDatum(m, v218)
								mBase = m.M
								v221 = m.ExcPending
								if v221 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(112)
									return v220
								}
							}
						} else {
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
							v73 = F_get_base_element_type(m, v72)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if v73 != v68 {
									if v66 == int32(2750) {
										v190 = float64(0.01)
									} else {
										v190 = float64(0.005)
									}
									v198 = v190
									v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
									if v199 != 0 {
										v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
										m.T0[v200].(func(*base.Module, int32))(m, v199)
										mBase = m.M
										v202 = m.ExcPending
										if v202 != 0 {
											return int32(0)
										} else {
											v203 = float64(0)
											if base.F64_lt(v198, v203) != 0 {
												v218 = v203
											} else {
												if base.F64_gt(v198, float64(1)) == int32(0) {
													v218 = v198
												} else {
													v218 = float64(1)
												}
											}
											v220 = F_Float8GetDatum(m, v218)
											mBase = m.M
											v221 = m.ExcPending
											if v221 != 0 {
												return int32(0)
											} else {
												m.G0 = v13 + int32(112)
												return v220
											}
										}
									} else {
										v203 = float64(0)
										if base.F64_lt(v198, v203) != 0 {
											v218 = v203
										} else {
											if base.F64_gt(v198, float64(1)) == int32(0) {
												v218 = v198
											} else {
												v218 = float64(1)
											}
										}
										v220 = F_Float8GetDatum(m, v218)
										mBase = m.M
										v221 = m.ExcPending
										if v221 != 0 {
											return int32(0)
										} else {
											m.G0 = v13 + int32(112)
											return v220
										}
									}
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
									v79 = F_lookup_type_cache(m, v68, int32(64))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+108))
										if v81 == int32(0) {
											if v66 == int32(2750) {
												v88 = float64(0.01)
											} else {
												v88 = float64(0.005)
											}
											v198 = v88
											v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
											if v199 != 0 {
												v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
												m.T0[v200].(func(*base.Module, int32))(m, v199)
												mBase = m.M
												v202 = m.ExcPending
												if v202 != 0 {
													return int32(0)
												} else {
													v203 = float64(0)
													if base.F64_lt(v198, v203) != 0 {
														v218 = v203
													} else {
														if base.F64_gt(v198, float64(1)) == int32(0) {
															v218 = v198
														} else {
															v218 = float64(1)
														}
													}
													v220 = F_Float8GetDatum(m, v218)
													mBase = m.M
													v221 = m.ExcPending
													if v221 != 0 {
														return int32(0)
													} else {
														m.G0 = v13 + int32(112)
														return v220
													}
												}
											} else {
												v203 = float64(0)
												if base.F64_lt(v198, v203) != 0 {
													v218 = v203
												} else {
													if base.F64_gt(v198, float64(1)) == int32(0) {
														v218 = v198
													} else {
														v218 = float64(1)
													}
												}
												v220 = F_Float8GetDatum(m, v218)
												mBase = m.M
												v221 = m.ExcPending
												if v221 != 0 {
													return int32(0)
												} else {
													m.G0 = v13 + int32(112)
													return v220
												}
											}
										} else {
											v89 = F_pg_detoast_datum(m, v77)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int32(0)
											} else {
												v91 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
												if v91 == int32(0) {
													v170 = int32(0)
													v176 = F_mcelem_array_selec(m, v89, v79, v170, v170, v170, v170, v170, v170, v66)
													mBase = m.M
													v177 = m.ExcPending
													if v177 != 0 {
														return int32(0)
													} else {
														v182 = v176
														if v89 == v77 {
															v198 = v182
															v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
															if v199 != 0 {
																v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																m.T0[v200].(func(*base.Module, int32))(m, v199)
																mBase = m.M
																v202 = m.ExcPending
																if v202 != 0 {
																	return int32(0)
																} else {
																	v203 = float64(0)
																	if base.F64_lt(v198, v203) != 0 {
																		v218 = v203
																	} else {
																		if base.F64_gt(v198, float64(1)) == int32(0) {
																			v218 = v198
																		} else {
																			v218 = float64(1)
																		}
																	}
																	v220 = F_Float8GetDatum(m, v218)
																	mBase = m.M
																	v221 = m.ExcPending
																	if v221 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v13 + int32(112)
																		return v220
																	}
																}
															} else {
																v203 = float64(0)
																if base.F64_lt(v198, v203) != 0 {
																	v218 = v203
																} else {
																	if base.F64_gt(v198, float64(1)) == int32(0) {
																		v218 = v198
																	} else {
																		v218 = float64(1)
																	}
																}
																v220 = F_Float8GetDatum(m, v218)
																mBase = m.M
																v221 = m.ExcPending
																if v221 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(112)
																	return v220
																}
															}
														} else {
															F_pfree(m, v89)
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return int32(0)
															} else {
																v198 = v182
																v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																if v199 != 0 {
																	v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																	m.T0[v200].(func(*base.Module, int32))(m, v199)
																	mBase = m.M
																	v202 = m.ExcPending
																	if v202 != 0 {
																		return int32(0)
																	} else {
																		v203 = float64(0)
																		if base.F64_lt(v198, v203) != 0 {
																			v218 = v203
																		} else {
																			if base.F64_gt(v198, float64(1)) == int32(0) {
																				v218 = v198
																			} else {
																				v218 = float64(1)
																			}
																		}
																		v220 = F_Float8GetDatum(m, v218)
																		mBase = m.M
																		v221 = m.ExcPending
																		if v221 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v13 + int32(112)
																			return v220
																		}
																	}
																} else {
																	v203 = float64(0)
																	if base.F64_lt(v198, v203) != 0 {
																		v218 = v203
																	} else {
																		if base.F64_gt(v198, float64(1)) == int32(0) {
																			v218 = v198
																		} else {
																			v218 = float64(1)
																		}
																	}
																	v220 = F_Float8GetDatum(m, v218)
																	mBase = m.M
																	v221 = m.ExcPending
																	if v221 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v13 + int32(112)
																		return v220
																	}
																}
															}
														}
													}
												} else {
													v96 = *(*int32)(unsafe.Add(mBase, uint32(v79)+108))
													v97 = F_statistic_proc_security_check(m, v13+int32(8), v96)
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return int32(0)
													} else {
														if v97 == int32(0) {
															v170 = int32(0)
															v176 = F_mcelem_array_selec(m, v89, v79, v170, v170, v170, v170, v170, v170, v66)
															mBase = m.M
															v177 = m.ExcPending
															if v177 != 0 {
																return int32(0)
															} else {
																v182 = v176
																if v89 == v77 {
																	v198 = v182
																	v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																	if v199 != 0 {
																		v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																		m.T0[v200].(func(*base.Module, int32))(m, v199)
																		mBase = m.M
																		v202 = m.ExcPending
																		if v202 != 0 {
																			return int32(0)
																		} else {
																			v203 = float64(0)
																			if base.F64_lt(v198, v203) != 0 {
																				v218 = v203
																			} else {
																				if base.F64_gt(v198, float64(1)) == int32(0) {
																					v218 = v198
																				} else {
																					v218 = float64(1)
																				}
																			}
																			v220 = F_Float8GetDatum(m, v218)
																			mBase = m.M
																			v221 = m.ExcPending
																			if v221 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v13 + int32(112)
																				return v220
																			}
																		}
																	} else {
																		v203 = float64(0)
																		if base.F64_lt(v198, v203) != 0 {
																			v218 = v203
																		} else {
																			if base.F64_gt(v198, float64(1)) == int32(0) {
																				v218 = v198
																			} else {
																				v218 = float64(1)
																			}
																		}
																		v220 = F_Float8GetDatum(m, v218)
																		mBase = m.M
																		v221 = m.ExcPending
																		if v221 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v13 + int32(112)
																			return v220
																		}
																	}
																} else {
																	F_pfree(m, v89)
																	mBase = m.M
																	v185 = m.ExcPending
																	if v185 != 0 {
																		return int32(0)
																	} else {
																		v198 = v182
																		v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																		if v199 != 0 {
																			v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																			m.T0[v200].(func(*base.Module, int32))(m, v199)
																			mBase = m.M
																			v202 = m.ExcPending
																			if v202 != 0 {
																				return int32(0)
																			} else {
																				v203 = float64(0)
																				if base.F64_lt(v198, v203) != 0 {
																					v218 = v203
																				} else {
																					if base.F64_gt(v198, float64(1)) == int32(0) {
																						v218 = v198
																					} else {
																						v218 = float64(1)
																					}
																				}
																				v220 = F_Float8GetDatum(m, v218)
																				mBase = m.M
																				v221 = m.ExcPending
																				if v221 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v13 + int32(112)
																					return v220
																				}
																			}
																		} else {
																			v203 = float64(0)
																			if base.F64_lt(v198, v203) != 0 {
																				v218 = v203
																			} else {
																				if base.F64_gt(v198, float64(1)) == int32(0) {
																					v218 = v198
																				} else {
																					v218 = float64(1)
																				}
																			}
																			v220 = F_Float8GetDatum(m, v218)
																			mBase = m.M
																			v221 = m.ExcPending
																			if v221 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v13 + int32(112)
																				return v220
																			}
																		}
																	}
																}
															}
														} else {
															v101 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
															v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
															v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+22)))
															v110 = F_get_attstatsslot(m, v13+int32(76), v101, int32(4), int32(0), int32(3))
															mBase = m.M
															v111 = m.ExcPending
															if v111 != 0 {
																return int32(0)
															} else {
																if v110 != 0 {
																	if v66 != int32(2752) {
																		v126 = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v126
																		v129 = int64(0)
																		*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v129
																		*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = v129
																		*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v129
																		*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v129
																		v138 = v126
																		v139 = v126
																		v140 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
																		v141 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
																		v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
																		v143 = *(*int32)(unsafe.Add(mBase, uint32(v13)+100))
																		v144 = F_mcelem_array_selec(m, v89, v79, v140, v141, v142, v143, v139, v138, v66)
																		mBase = m.M
																		v145 = m.ExcPending
																		if v145 != 0 {
																			return int32(0)
																		} else {
																			F_free_attstatsslot(m, v13+int32(40))
																			mBase = m.M
																			v149 = m.ExcPending
																			if v149 != 0 {
																				return int32(0)
																			} else {
																				F_free_attstatsslot(m, v13+int32(76))
																				mBase = m.M
																				v153 = m.ExcPending
																				if v153 != 0 {
																					return int32(0)
																				} else {
																					v164 = v144
																					v166 = *(*float32)(unsafe.Add(mBase, uint32(v102+v103)+8))
																					v182 = base.F64_mul(v164, base.F64_sub(float64(1), base.F64_promote_f32(v166)))
																					if v89 == v77 {
																						v198 = v182
																						v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																						if v199 != 0 {
																							v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																							m.T0[v200].(func(*base.Module, int32))(m, v199)
																							mBase = m.M
																							v202 = m.ExcPending
																							if v202 != 0 {
																								return int32(0)
																							} else {
																								v203 = float64(0)
																								if base.F64_lt(v198, v203) != 0 {
																									v218 = v203
																								} else {
																									if base.F64_gt(v198, float64(1)) == int32(0) {
																										v218 = v198
																									} else {
																										v218 = float64(1)
																									}
																								}
																								v220 = F_Float8GetDatum(m, v218)
																								mBase = m.M
																								v221 = m.ExcPending
																								if v221 != 0 {
																									return int32(0)
																								} else {
																									m.G0 = v13 + int32(112)
																									return v220
																								}
																							}
																						} else {
																							v203 = float64(0)
																							if base.F64_lt(v198, v203) != 0 {
																								v218 = v203
																							} else {
																								if base.F64_gt(v198, float64(1)) == int32(0) {
																									v218 = v198
																								} else {
																									v218 = float64(1)
																								}
																							}
																							v220 = F_Float8GetDatum(m, v218)
																							mBase = m.M
																							v221 = m.ExcPending
																							if v221 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v13 + int32(112)
																								return v220
																							}
																						}
																					} else {
																						F_pfree(m, v89)
																						mBase = m.M
																						v185 = m.ExcPending
																						if v185 != 0 {
																							return int32(0)
																						} else {
																							v198 = v182
																							v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																							if v199 != 0 {
																								v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																								m.T0[v200].(func(*base.Module, int32))(m, v199)
																								mBase = m.M
																								v202 = m.ExcPending
																								if v202 != 0 {
																									return int32(0)
																								} else {
																									v203 = float64(0)
																									if base.F64_lt(v198, v203) != 0 {
																										v218 = v203
																									} else {
																										if base.F64_gt(v198, float64(1)) == int32(0) {
																											v218 = v198
																										} else {
																											v218 = float64(1)
																										}
																									}
																									v220 = F_Float8GetDatum(m, v218)
																									mBase = m.M
																									v221 = m.ExcPending
																									if v221 != 0 {
																										return int32(0)
																									} else {
																										m.G0 = v13 + int32(112)
																										return v220
																									}
																								}
																							} else {
																								v203 = float64(0)
																								if base.F64_lt(v198, v203) != 0 {
																									v218 = v203
																								} else {
																									if base.F64_gt(v198, float64(1)) == int32(0) {
																										v218 = v198
																									} else {
																										v218 = float64(1)
																									}
																								}
																								v220 = F_Float8GetDatum(m, v218)
																								mBase = m.M
																								v221 = m.ExcPending
																								if v221 != 0 {
																									return int32(0)
																								} else {
																									m.G0 = v13 + int32(112)
																									return v220
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																		v120 = F_get_attstatsslot(m, v13+int32(40), v116, int32(5), int32(0), int32(2))
																		mBase = m.M
																		v121 = m.ExcPending
																		if v121 != 0 {
																			return int32(0)
																		} else {
																			if v120 == int32(0) {
																				v126 = int32(0)
																				*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v126
																				v129 = int64(0)
																				*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v129
																				*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = v129
																				*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v129
																				*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v129
																				v138 = v126
																				v139 = v126
																			} else {
																				v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
																				v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
																				v138 = v124
																				v139 = v125
																			}
																			v140 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
																			v141 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
																			v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
																			v143 = *(*int32)(unsafe.Add(mBase, uint32(v13)+100))
																			v144 = F_mcelem_array_selec(m, v89, v79, v140, v141, v142, v143, v139, v138, v66)
																			mBase = m.M
																			v145 = m.ExcPending
																			if v145 != 0 {
																				return int32(0)
																			} else {
																				F_free_attstatsslot(m, v13+int32(40))
																				mBase = m.M
																				v149 = m.ExcPending
																				if v149 != 0 {
																					return int32(0)
																				} else {
																					F_free_attstatsslot(m, v13+int32(76))
																					mBase = m.M
																					v153 = m.ExcPending
																					if v153 != 0 {
																						return int32(0)
																					} else {
																						v164 = v144
																						v166 = *(*float32)(unsafe.Add(mBase, uint32(v102+v103)+8))
																						v182 = base.F64_mul(v164, base.F64_sub(float64(1), base.F64_promote_f32(v166)))
																						if v89 == v77 {
																							v198 = v182
																							v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																							if v199 != 0 {
																								v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																								m.T0[v200].(func(*base.Module, int32))(m, v199)
																								mBase = m.M
																								v202 = m.ExcPending
																								if v202 != 0 {
																									return int32(0)
																								} else {
																									v203 = float64(0)
																									if base.F64_lt(v198, v203) != 0 {
																										v218 = v203
																									} else {
																										if base.F64_gt(v198, float64(1)) == int32(0) {
																											v218 = v198
																										} else {
																											v218 = float64(1)
																										}
																									}
																									v220 = F_Float8GetDatum(m, v218)
																									mBase = m.M
																									v221 = m.ExcPending
																									if v221 != 0 {
																										return int32(0)
																									} else {
																										m.G0 = v13 + int32(112)
																										return v220
																									}
																								}
																							} else {
																								v203 = float64(0)
																								if base.F64_lt(v198, v203) != 0 {
																									v218 = v203
																								} else {
																									if base.F64_gt(v198, float64(1)) == int32(0) {
																										v218 = v198
																									} else {
																										v218 = float64(1)
																									}
																								}
																								v220 = F_Float8GetDatum(m, v218)
																								mBase = m.M
																								v221 = m.ExcPending
																								if v221 != 0 {
																									return int32(0)
																								} else {
																									m.G0 = v13 + int32(112)
																									return v220
																								}
																							}
																						} else {
																							F_pfree(m, v89)
																							mBase = m.M
																							v185 = m.ExcPending
																							if v185 != 0 {
																								return int32(0)
																							} else {
																								v198 = v182
																								v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																								if v199 != 0 {
																									v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																									m.T0[v200].(func(*base.Module, int32))(m, v199)
																									mBase = m.M
																									v202 = m.ExcPending
																									if v202 != 0 {
																										return int32(0)
																									} else {
																										v203 = float64(0)
																										if base.F64_lt(v198, v203) != 0 {
																											v218 = v203
																										} else {
																											if base.F64_gt(v198, float64(1)) == int32(0) {
																												v218 = v198
																											} else {
																												v218 = float64(1)
																											}
																										}
																										v220 = F_Float8GetDatum(m, v218)
																										mBase = m.M
																										v221 = m.ExcPending
																										if v221 != 0 {
																											return int32(0)
																										} else {
																											m.G0 = v13 + int32(112)
																											return v220
																										}
																									}
																								} else {
																									v203 = float64(0)
																									if base.F64_lt(v198, v203) != 0 {
																										v218 = v203
																									} else {
																										if base.F64_gt(v198, float64(1)) == int32(0) {
																											v218 = v198
																										} else {
																											v218 = float64(1)
																										}
																									}
																									v220 = F_Float8GetDatum(m, v218)
																									mBase = m.M
																									v221 = m.ExcPending
																									if v221 != 0 {
																										return int32(0)
																									} else {
																										m.G0 = v13 + int32(112)
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
																} else {
																	v154 = int32(0)
																	v160 = F_mcelem_array_selec(m, v89, v79, v154, v154, v154, v154, v154, v154, v66)
																	mBase = m.M
																	v161 = m.ExcPending
																	if v161 != 0 {
																		return int32(0)
																	} else {
																		v164 = v160
																		v166 = *(*float32)(unsafe.Add(mBase, uint32(v102+v103)+8))
																		v182 = base.F64_mul(v164, base.F64_sub(float64(1), base.F64_promote_f32(v166)))
																		if v89 == v77 {
																			v198 = v182
																			v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																			if v199 != 0 {
																				v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																				m.T0[v200].(func(*base.Module, int32))(m, v199)
																				mBase = m.M
																				v202 = m.ExcPending
																				if v202 != 0 {
																					return int32(0)
																				} else {
																					v203 = float64(0)
																					if base.F64_lt(v198, v203) != 0 {
																						v218 = v203
																					} else {
																						if base.F64_gt(v198, float64(1)) == int32(0) {
																							v218 = v198
																						} else {
																							v218 = float64(1)
																						}
																					}
																					v220 = F_Float8GetDatum(m, v218)
																					mBase = m.M
																					v221 = m.ExcPending
																					if v221 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v13 + int32(112)
																						return v220
																					}
																				}
																			} else {
																				v203 = float64(0)
																				if base.F64_lt(v198, v203) != 0 {
																					v218 = v203
																				} else {
																					if base.F64_gt(v198, float64(1)) == int32(0) {
																						v218 = v198
																					} else {
																						v218 = float64(1)
																					}
																				}
																				v220 = F_Float8GetDatum(m, v218)
																				mBase = m.M
																				v221 = m.ExcPending
																				if v221 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v13 + int32(112)
																					return v220
																				}
																			}
																		} else {
																			F_pfree(m, v89)
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
																				return int32(0)
																			} else {
																				v198 = v182
																				v199 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																				if v199 != 0 {
																					v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																					m.T0[v200].(func(*base.Module, int32))(m, v199)
																					mBase = m.M
																					v202 = m.ExcPending
																					if v202 != 0 {
																						return int32(0)
																					} else {
																						v203 = float64(0)
																						if base.F64_lt(v198, v203) != 0 {
																							v218 = v203
																						} else {
																							if base.F64_gt(v198, float64(1)) == int32(0) {
																								v218 = v198
																							} else {
																								v218 = float64(1)
																							}
																						}
																						v220 = F_Float8GetDatum(m, v218)
																						mBase = m.M
																						v221 = m.ExcPending
																						if v221 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v13 + int32(112)
																							return v220
																						}
																					}
																				} else {
																					v203 = float64(0)
																					if base.F64_lt(v198, v203) != 0 {
																						v218 = v203
																					} else {
																						if base.F64_gt(v198, float64(1)) == int32(0) {
																							v218 = v198
																						} else {
																							v218 = float64(1)
																						}
																					}
																					v220 = F_Float8GetDatum(m, v218)
																					mBase = m.M
																					v221 = m.ExcPending
																					if v221 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v13 + int32(112)
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
	v754 = F_expression_tree_walker_impl(m, l0, int32(483), v15+int32(48))
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
	v257 = F_expression_tree_walker_impl(m, l0, int32(483), v15+int32(48))
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
	v240 = F_expression_tree_walker_impl(m, l0, int32(483), v15+int32(48))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L21
	} else {
		goto L77
	}
L15:
	;
	v193 = F_expression_tree_walker_impl(m, l0, int32(483), v15+int32(48))
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
	v44 = F_expression_tree_walker_impl(m, l0, int32(483), v15+int32(48))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L21
	} else {
		goto L23
	}
L20:
	;
	v34 = F_expression_tree_walker_impl(m, l0, int32(483), v15+int32(48))
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
	v170 = F_expression_tree_walker_impl(m, l0, int32(483), v15+int32(48))
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
					F_errfinish(m, int32(_a_F_attnumAttName_1), int32(3649), int32(_a_F_attnumAttName_2))
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
			v29 = v16 + v17<<(uint(int32(4))%32) + l1*int32(100) - int32(80)
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
					F_errfinish(m, int32(_a_F_attnumCollationId_1), int32(3689), int32(_a_F_attnumCollationId_2))
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
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13<<(uint(int32(4))%32)+l1*int32(100))+16))
			v24 = v21
			m.G0 = v7 + int32(16)
			return v24
		}
	}
}
