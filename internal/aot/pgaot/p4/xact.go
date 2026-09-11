package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_xactGetCommittedInvalidationMessages(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_xactGetCommittedInvalidationMessages[0]))
	if v7 == v3 {
		v10 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v10)
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v10
		return v10
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)))
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v16)
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_xactGetCommittedInvalidationMessages[1]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
		v37 = F_MemoryContextAlloc(m, v19, (v20+(v21+(v22+v23))-(v27+(v28+(v29+v30))))<<(uint(int32(4))%32))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v37
			v43 = *(*int32)(unsafe.Add(mBase, _c_F_xactGetCommittedInvalidationMessages[0]))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+28))
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
			v46 = v44 - v45
			if int32(0) < v46 {
				v50 = *(*int32)(unsafe.Add(mBase, _c_F_xactGetCommittedInvalidationMessages[2]))
				v51 = int32(4)
				v55 = v46 << (uint(v51) % 32)
				if v55 != 0 {
					v56 = F__emscripten_memcpy_bulkmem(m, v37, v50+v45<<(uint(v51)%32), v55)
					mBase = m.M
				} else {
				}
				v58 = v46
			} else {
				v58 = v3
			}
			v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
			v61 = v59 - v60
			if int32(0) < v61 {
				v64 = int32(4)
				v68 = *(*int32)(unsafe.Add(mBase, _c_F_xactGetCommittedInvalidationMessages[2]))
				v73 = v61 << (uint(v64) % 32)
				if v73 != 0 {
					v74 = F__emscripten_memcpy_bulkmem(m, v37+v58<<(uint(v64)%32), v68+v60<<(uint(v64)%32), v73)
					mBase = m.M
				} else {
				}
				v77 = v61 + v58
			} else {
				v77 = v58
			}
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v43)+32))
			v79 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
			v80 = v78 - v79
			if int32(0) < v80 {
				v83 = int32(4)
				v87 = *(*int32)(unsafe.Add(mBase, _c_F_xactGetCommittedInvalidationMessages[3]))
				v92 = v80 << (uint(v83) % 32)
				if v92 != 0 {
					v93 = F__emscripten_memcpy_bulkmem(m, v37+v77<<(uint(v83)%32), v87+v79<<(uint(v83)%32), v92)
					mBase = m.M
				} else {
				}
				v96 = v80 + v77
			} else {
				v96 = v77
			}
			v97 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
			v99 = v97 - v98
			if int32(0) < v99 {
				v102 = int32(4)
				v106 = *(*int32)(unsafe.Add(mBase, _c_F_xactGetCommittedInvalidationMessages[3]))
				v111 = v99 << (uint(v102) % 32)
				if v111 != 0 {
					v112 = F__emscripten_memcpy_bulkmem(m, v37+v96<<(uint(v102)%32), v106+v98<<(uint(v102)%32), v111)
					mBase = m.M
				} else {
				}
				v115 = v99 + v96
			} else {
				v115 = v96
			}
			return v115
		}
	}
}
func F_xact_redo_abort(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int64
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = v6 - int32(1)
	if v12 < int32(0) {
		v79 = l1
	} else {
		if v6&int32(1) != 0 {
			v17 = int32(2)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v7+v12<<(uint(v17)%32))))
			if base.B2i32(base.Ui32(v17) < base.Ui32(v20))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l1)) == int32(0) {
				if base.Ui32(l1) < base.Ui32(v20) {
					v32 = v20
				} else {
					v32 = l1
				}
			} else {
				if int32(0) <= l1-v20 {
					v32 = l1
				} else {
					v32 = v20
				}
			}
			v35 = v32
			v37 = v6 - int32(2)
		} else {
			v35 = l1
			v37 = v12
		}
		if v12 == int32(0) {
			v79 = v35
		} else {
			v42 = v35
			v43 = v37
			for {
				v49 = v43 << (uint(int32(2)) % 32)
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v7+v49)))
				if base.Ui32(v42) < base.Ui32(int32(3)) {
					if base.Ui32(v51) <= base.Ui32(v42) {
						v60 = v42
					} else {
						v60 = v51
					}
				} else {
					if base.Ui32(v51) < base.Ui32(int32(3)) {
						if base.Ui32(v51) <= base.Ui32(v42) {
							v60 = v42
						} else {
							v60 = v51
						}
					} else {
						if v42-v51 < int32(0) {
							v60 = v51
						} else {
							v60 = v42
						}
					}
				}
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v7-int32(4)))))
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v62))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v60)) == int32(0) {
					if base.Ui32(v60) < base.Ui32(v62) {
						v74 = v62
					} else {
						v74 = v60
					}
				} else {
					if int32(0) <= v60-v62 {
						v74 = v60
					} else {
						v74 = v62
					}
				}
				if int32(1) < v43 {
					v42 = v74
					v43 = v43 - int32(2)
					continue
				} else {
					break
				}
				break
			}
			v79 = v74
		}
	}
	F_AdvanceNextFullTransactionIdPastXid(m, v79)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		return
	} else {
		v88 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo_abort[0]))
		if v88 == int32(0) {
			v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			F_TransactionIdAbortTree(m, l1, v91, v92)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return
			} else {
				v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
				if v114&int32(32) != 0 {
					v117 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
					v118 = int32(0)
					F_replorigin_advance(m, l3, v117, l2, v118, v118)
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return
					} else {
						v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if int32(0) < v122 {
							F_XLogFlush(m, l2)
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return
							} else {
								v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								F_DropRelationFiles(m, v127, v128, int32(1))
								mBase = m.M
								v131 = m.ExcPending
								if v131 != 0 {
									return
								} else {
									v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if int32(0) < v132 {
										F_XLogFlush(m, l2)
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return
										} else {
											v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											F_pgstat_execute_transactional_drops(m, v137, v138)
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return
											} else {
												return
											}
										}
									} else {
										return
									}
								}
							}
						} else {
							v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if int32(0) < v132 {
								F_XLogFlush(m, l2)
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
									return
								} else {
									v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									F_pgstat_execute_transactional_drops(m, v137, v138)
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								return
							}
						}
					}
				} else {
					v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if int32(0) < v122 {
						F_XLogFlush(m, l2)
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return
						} else {
							v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							F_DropRelationFiles(m, v127, v128, int32(1))
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return
							} else {
								v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if int32(0) < v132 {
									F_XLogFlush(m, l2)
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return
									} else {
										v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										F_pgstat_execute_transactional_drops(m, v137, v138)
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return
										} else {
											return
										}
									}
								} else {
									return
								}
							}
						}
					} else {
						v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if int32(0) < v132 {
							F_XLogFlush(m, l2)
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return
							} else {
								v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								F_pgstat_execute_transactional_drops(m, v137, v138)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							return
						}
					}
				}
			}
		} else {
			F_RecordKnownAssignedTransactionIds(m, v79)
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return
			} else {
				v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				F_TransactionIdAbortTree(m, l1, v97, v98)
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return
				} else {
					v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					F_ExpireTreeKnownAssignedTransactionIds(m, l1, v101, v102, v79)
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return
					} else {
						v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
						if v105&int32(64) == int32(0) {
							v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
							if v114&int32(32) != 0 {
								v117 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
								v118 = int32(0)
								F_replorigin_advance(m, l3, v117, l2, v118, v118)
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return
								} else {
									v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									if int32(0) < v122 {
										F_XLogFlush(m, l2)
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return
										} else {
											v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											F_DropRelationFiles(m, v127, v128, int32(1))
											mBase = m.M
											v131 = m.ExcPending
											if v131 != 0 {
												return
											} else {
												v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if int32(0) < v132 {
													F_XLogFlush(m, l2)
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return
													} else {
														v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														F_pgstat_execute_transactional_drops(m, v137, v138)
														mBase = m.M
														v140 = m.ExcPending
														if v140 != 0 {
															return
														} else {
															return
														}
													}
												} else {
													return
												}
											}
										}
									} else {
										v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if int32(0) < v132 {
											F_XLogFlush(m, l2)
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return
											} else {
												v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
												F_pgstat_execute_transactional_drops(m, v137, v138)
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return
												} else {
													return
												}
											}
										} else {
											return
										}
									}
								}
							} else {
								v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if int32(0) < v122 {
									F_XLogFlush(m, l2)
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return
									} else {
										v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										F_DropRelationFiles(m, v127, v128, int32(1))
										mBase = m.M
										v131 = m.ExcPending
										if v131 != 0 {
											return
										} else {
											v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if int32(0) < v132 {
												F_XLogFlush(m, l2)
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return
												} else {
													v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
													F_pgstat_execute_transactional_drops(m, v137, v138)
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return
													} else {
														return
													}
												}
											} else {
												return
											}
										}
									}
								} else {
									v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if int32(0) < v132 {
										F_XLogFlush(m, l2)
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return
										} else {
											v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											F_pgstat_execute_transactional_drops(m, v137, v138)
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return
											} else {
												return
											}
										}
									} else {
										return
									}
								}
							}
						} else {
							v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							F_StandbyReleaseLockTree(m, l1, v110, v111)
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return
							} else {
								v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
								if v114&int32(32) != 0 {
									v117 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
									v118 = int32(0)
									F_replorigin_advance(m, l3, v117, l2, v118, v118)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return
									} else {
										v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										if int32(0) < v122 {
											F_XLogFlush(m, l2)
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
												return
											} else {
												v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												F_DropRelationFiles(m, v127, v128, int32(1))
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return
												} else {
													v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if int32(0) < v132 {
														F_XLogFlush(m, l2)
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
															return
														} else {
															v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
															v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
															F_pgstat_execute_transactional_drops(m, v137, v138)
															mBase = m.M
															v140 = m.ExcPending
															if v140 != 0 {
																return
															} else {
																return
															}
														}
													} else {
														return
													}
												}
											}
										} else {
											v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if int32(0) < v132 {
												F_XLogFlush(m, l2)
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return
												} else {
													v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
													F_pgstat_execute_transactional_drops(m, v137, v138)
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return
													} else {
														return
													}
												}
											} else {
												return
											}
										}
									}
								} else {
									v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									if int32(0) < v122 {
										F_XLogFlush(m, l2)
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return
										} else {
											v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											F_DropRelationFiles(m, v127, v128, int32(1))
											mBase = m.M
											v131 = m.ExcPending
											if v131 != 0 {
												return
											} else {
												v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if int32(0) < v132 {
													F_XLogFlush(m, l2)
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return
													} else {
														v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														F_pgstat_execute_transactional_drops(m, v137, v138)
														mBase = m.M
														v140 = m.ExcPending
														if v140 != 0 {
															return
														} else {
															return
														}
													}
												} else {
													return
												}
											}
										}
									} else {
										v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if int32(0) < v132 {
											F_XLogFlush(m, l2)
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return
											} else {
												v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
												F_pgstat_execute_transactional_drops(m, v137, v138)
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return
												} else {
													return
												}
											}
										} else {
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
	}
}
