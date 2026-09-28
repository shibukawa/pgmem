package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_performDeletion(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v68 int32
	_ = v68
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v14
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v17 == int32(1259) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			F_LockRelationOid(m, v20, int32(8))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v109 = F_palloc(m, int32(16))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v109)+8)) = int64(137438953472)
					v115 = F_palloc_mul(m, int32(12), int32(32))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						v117 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v117
						*(*int32)(unsafe.Add(mBase, uint32(v109))) = v115
						v124 = v10 + int32(12)
						F_findDependentObjects(m, l0, int32(1), l2, v117, v109, v117, v124)
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return
						} else {
							F_reportDependentObjects(m, v109, l1, l2, l0)
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return
							} else {
								F_deleteObjectsInList(m, v109, v124, l2)
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return
								} else {
									v131 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
									F_pfree(m, v131)
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return
									} else {
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
										if v134 != 0 {
											F_pfree(m, v134)
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return
											} else {
												F_pfree(m, v109)
												mBase = m.M
												v138 = m.ExcPending
												if v138 != 0 {
													return
												} else {
													v139 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
													F_relation_close(m, v139, int32(3))
													mBase = m.M
													v142 = m.ExcPending
													if v142 != 0 {
														return
													} else {
														m.G0 = v10 + int32(16)
														return
													}
												}
											}
										} else {
											F_pfree(m, v109)
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return
											} else {
												v139 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												F_relation_close(m, v139, int32(3))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return
												} else {
													m.G0 = v10 + int32(16)
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
		} else {
			v26 = int32(1)
			if v17 <= int32(3591) {
				if v17 <= int32(2670) {
					switch v17 - int32(1213) {
					case 0, 1, 19, 20, 47, 48, 49:
						v97 = v26
					case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
						v97 = int32(0)
					default:
						if base.Ui32(int32(2)) <= base.Ui32(v17-int32(2396)) {
							v97 = int32(0)
						} else {
							v97 = v26
						}
					}
				} else {
					v38 = v17 - int32(2671)
					if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v38))|base.B2i32(int32(1)<<(uint(v38)%32)&int32(226492515) == int32(0)) != 0 {
						if base.B2i32(base.Ui32(v17-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v17-int32(2846)) < base.Ui32(int32(2))) != 0 {
							v97 = v26
						} else {
							v97 = int32(0)
						}
					} else {
						v97 = v26
					}
				}
			} else {
				if v17 <= int32(_a_F_performDeletion_0) {
					v51 = v17 - int32(_a_F_performDeletion_1)
					if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v51))|base.B2i32(int32(1)<<(uint(v51)%32)&int32(963) == int32(0)) != 0 {
						if base.Ui32(v17-int32(3592)) < base.Ui32(int32(2)) {
							v97 = v26
						} else {
							if base.Ui32(int32(2)) <= base.Ui32(v17-int32(4060)) {
								v97 = int32(0)
							} else {
								v97 = v26
							}
						}
					} else {
						v97 = v26
					}
				} else {
					switch v17 - int32(_a_F_performDeletion_2) {
					case 0, 1, 2, 3, 4, 59, 60:
						v97 = v26
					case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
						v97 = int32(0)
					default:
						if base.Ui32(v17-int32(_a_F_performDeletion_3)) < base.Ui32(int32(3)) {
							v97 = v26
						} else {
							v68 = v17 - int32(_a_F_performDeletion_4)
							if base.Ui32(int32(15)) < base.Ui32(v68) {
								v97 = int32(0)
							} else {
								if int32(1)<<(uint(v68)%32)&int32(_a_F_performDeletion_5) != 0 {
									v97 = v26
								} else {
									v97 = int32(0)
								}
							}
						}
					}
				}
			}
			v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v97 != 0 {
				F_LockSharedObject(m, v99, v98, int32(8))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return
				} else {
					v109 = F_palloc(m, int32(16))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v109)+8)) = int64(137438953472)
						v115 = F_palloc_mul(m, int32(12), int32(32))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							v117 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v117
							*(*int32)(unsafe.Add(mBase, uint32(v109))) = v115
							v124 = v10 + int32(12)
							F_findDependentObjects(m, l0, int32(1), l2, v117, v109, v117, v124)
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return
							} else {
								F_reportDependentObjects(m, v109, l1, l2, l0)
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return
								} else {
									F_deleteObjectsInList(m, v109, v124, l2)
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return
									} else {
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
										F_pfree(m, v131)
										mBase = m.M
										v133 = m.ExcPending
										if v133 != 0 {
											return
										} else {
											v134 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
											if v134 != 0 {
												F_pfree(m, v134)
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return
												} else {
													F_pfree(m, v109)
													mBase = m.M
													v138 = m.ExcPending
													if v138 != 0 {
														return
													} else {
														v139 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
														F_relation_close(m, v139, int32(3))
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return
														} else {
															m.G0 = v10 + int32(16)
															return
														}
													}
												}
											} else {
												F_pfree(m, v109)
												mBase = m.M
												v138 = m.ExcPending
												if v138 != 0 {
													return
												} else {
													v139 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
													F_relation_close(m, v139, int32(3))
													mBase = m.M
													v142 = m.ExcPending
													if v142 != 0 {
														return
													} else {
														m.G0 = v10 + int32(16)
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
			} else {
				F_LockDatabaseObject(m, v99, v98, int32(8))
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return
				} else {
					v109 = F_palloc(m, int32(16))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v109)+8)) = int64(137438953472)
						v115 = F_palloc_mul(m, int32(12), int32(32))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							v117 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v117
							*(*int32)(unsafe.Add(mBase, uint32(v109))) = v115
							v124 = v10 + int32(12)
							F_findDependentObjects(m, l0, int32(1), l2, v117, v109, v117, v124)
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return
							} else {
								F_reportDependentObjects(m, v109, l1, l2, l0)
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return
								} else {
									F_deleteObjectsInList(m, v109, v124, l2)
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return
									} else {
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
										F_pfree(m, v131)
										mBase = m.M
										v133 = m.ExcPending
										if v133 != 0 {
											return
										} else {
											v134 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
											if v134 != 0 {
												F_pfree(m, v134)
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return
												} else {
													F_pfree(m, v109)
													mBase = m.M
													v138 = m.ExcPending
													if v138 != 0 {
														return
													} else {
														v139 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
														F_relation_close(m, v139, int32(3))
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return
														} else {
															m.G0 = v10 + int32(16)
															return
														}
													}
												}
											} else {
												F_pfree(m, v109)
												mBase = m.M
												v138 = m.ExcPending
												if v138 != 0 {
													return
												} else {
													v139 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
													F_relation_close(m, v139, int32(3))
													mBase = m.M
													v142 = m.ExcPending
													if v142 != 0 {
														return
													} else {
														m.G0 = v10 + int32(16)
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
		}
	}
}
